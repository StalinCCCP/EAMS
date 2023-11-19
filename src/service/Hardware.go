package service

import (
	"EAMSbackend/dbc"
	"EAMSbackend/define"
	"EAMSbackend/models"
	"EAMSbackend/util"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func HardwareCategoryQuery(c *gin.Context) {
	var data []util.NullString
	err := dbc.DB().Model(&models.Hardware{}).Distinct("Category").Pluck("Category", &data).Error
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

//	func HardwareStatusQuery(c *gin.Context) {
//		var data []util.NullString
//		err := dbc.DB().Model(&models.Hardware{}).Distinct("status").Pluck("status", &data).Error
//		if err != nil {
//			c.JSON(http.StatusInternalServerError, gin.H{
//				"msg": err,
//			})
//			return
//		}
//		c.JSON(http.StatusOK, gin.H{
//			"data": data,
//		})
//	}
func HardwareLocationQuery(c *gin.Context) {
	var data []util.NullString
	err := dbc.DB().Model(&models.Hardware{}).Distinct("location").Pluck("location", &data).Error
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
func HardwareListQuery(c *gin.Context) {
	name := define.HLQreq{}
	transfer, ok := c.Get("LocalCallData")
	if !ok {
		if err := c.ShouldBindJSON(&name); err != nil {
			log.Println("Bad request")
			c.Status(http.StatusBadRequest)
			return
		}
	} else {
		name = transfer.(define.HLQreq)
		delete(c.Keys, "LocalCallData")
	}
	query := dbc.DB().Model(&models.Hardware{})
	if name.HardwareName != "" {
		query = query.Where("hardware_name LIKE ?", "%"+name.HardwareName+"%")
	}
	if name.Category != "" {
		query = query.Where("Category = ?", name.Category)
	}
	if name.Status != "" {
		query = query.Where("Status = ?", name.Status)
	}
	if name.Location.String != "" {
		query = query.Where("Location = ?", name.Location)
	}
	var data []define.HLQresp
	err := query.Select("hardware_id, hardware_name,category,status,location").Scan(&data).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err,
		})
		return
	}
	// jsonData, err := json.Marshal(data)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"msg": err,
	// 	})
	// 	return
	// }
	//fmt.Println(string(jsonData))
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	} else {
		c.Set("LocalCallResult", data)
	}
}

func HardwareDetailQuery(c *gin.Context) {
	req := struct {
		HardwareID uint
	}{}
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("Bad request")
		c.Status(http.StatusBadRequest)
		return
	}
	var data1 []models.Hardware
	query := dbc.DB().Model(&models.Hardware{}).Where("hardware_id = ?", req.HardwareID)
	err := query.Find(&data1).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err,
		})
		return
	}
	var data2 []models.HardwareMaintenance
	query = dbc.DB().Model(&models.HardwareMaintenance{}).Where("hardware_id = ?", req.HardwareID)
	err = query.Find(&data2).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": map[string]interface{}{
			"Hinfo":     data1,
			"Maintinfo": data2,
		},
	})
}
func HardwareUpdate(c *gin.Context) { //TODO:对于没有填入的项目 如何区分呢
	req := models.Hardware{}
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("Bad request")
		c.Status(http.StatusBadRequest)
		return
	}
	set := map[string]bool{
		"占用":  true,
		"保留":  true,
		"正常":  true,
		"非正常": true,
	}
	if !set[req.Status] {
		log.Println("Bad request")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Status is not defined",
		})
		return
	}
	query := dbc.DB().Model(&models.Hardware{}).Where("hardware_id = ?", req.HardwareID)
	queryM := dbc.DB().Model(&models.HardwareMaintenance{}).Where("hardware_id = ? and status = '待处理'", req.HardwareID)
	chk := new(models.HardwareMaintenance)
	err := queryM.First(chk).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err,
		})
		return
	} else if req.Status != "非正常" && err == nil {
		c.JSON(http.StatusFailedDependency, gin.H{
			"msg": "Attempting to transform an abnormal hardware with maintenance processes unfinished to a normal one.",
		})
		return
	}
	if err := query.Updates(&req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err,
		})
		return
	}
	c.Status(http.StatusOK)
}

func HardwareDelete(c *gin.Context) {
	req := struct {
		HardwareID uint
	}{}
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("Bad request")
		c.Status(http.StatusBadRequest)
		return
	}
	query := dbc.DB().Model(&models.Hardware{}).Where("hardware_id = ?", req.HardwareID)
	queryM := dbc.DB().Model(&models.HardwareMaintenance{}).Where("hardware_id = ?", req.HardwareID)
	var tgt models.Hardware
	if err := query.Find(&tgt).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err,
		})
		return
	}
	chk := new(models.HardwareMaintenance)
	err := queryM.First(chk).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err,
		})
		return
	} else if err == nil {
		c.JSON(http.StatusFailedDependency, gin.H{
			"msg": "Attempting to remove a hardware with maintenance.",
		})
		return
	}
	if err := dbc.DB().Delete(&tgt).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err,
		})
		return
	}
	c.Status(http.StatusOK)
}

func HardwareCreate(c *gin.Context) {
	req := models.Hardware{}
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("Bad request")
		c.Status(http.StatusBadRequest)
		return
	}
	set := map[string]bool{
		"占用":  true,
		"保留":  true,
		"正常":  true,
		"非正常": true,
	}
	if !set[req.Status] {
		log.Println("Bad request")
		c.Status(http.StatusBadRequest)
		return
	}
	if err := dbc.DB().Create(&req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err,
		})
		return
	}
	c.Status(http.StatusOK)
}
