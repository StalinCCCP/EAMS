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

// func SoftwareVersionQuery(c *gin.Context) {
// 	var data []util.NullString
// 	err := dbc.DB().Model(&models.Software{}).Distinct("Version").Pluck("Version", &data).Error
// 	if err != nil {
// 		c.Status(http.StatusInternalServerError)
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"data": data,
// 	})
// }

//	func SoftwareStatusQuery(c *gin.Context) {
//		var data []util.NullString
//		err := dbc.DB().Model(&models.Software{}).Distinct("status").Pluck("status", &data).Error
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
func SoftwareLocationQuery(c *gin.Context) {
	var data []util.NullString
	err := dbc.DB().Model(&models.Software{}).Distinct("location").Pluck("location", &data).Error
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
func SoftwareListQuery(c *gin.Context) {
	name := define.SLQreq{}
	transfer, ok := c.Get("LocalCallData")
	if !ok {
		if err := c.ShouldBindJSON(&name); err != nil {
			log.Println("Bad request")
			c.Status(http.StatusBadRequest)
			return
		}
	} else {
		name = transfer.(define.SLQreq)
		delete(c.Keys, "LocalCallData")
	}
	query := dbc.DB().Model(&models.Software{})
	if name.SoftwareName != "" {
		query = query.Where("Software_name LIKE ?", "%"+name.SoftwareName+"%")
	}
	if name.Version != "" {
		query = query.Where("Version = ?", name.Version)
	}
	if name.Status != "" {
		query = query.Where("Status = ?", name.Status)
	}
	if name.Location.String != "" {
		query = query.Where("Location = ?", name.Location)
	}
	var data []define.SLQresp
	err := query.Select("Software_id, Software_name,Version,status,location").Scan(&data).Error
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

func SoftwareDetailQuery(c *gin.Context) {
	req := struct {
		SoftwareID uint
	}{}
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("Bad request")
		c.Status(http.StatusBadRequest)
		return
	}
	var data1 models.Software
	query := dbc.DB().Model(&models.Software{}).Where("Software_id = ?", req.SoftwareID)
	err := query.First(&data1).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err,
		})
		return
	}
	var data2 []models.SoftwareMaintenance
	query = dbc.DB().Model(&models.SoftwareMaintenance{}).Where("Software_id = ?", req.SoftwareID)
	err = query.Find(&data2).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": map[string]interface{}{
			"Sinfo":     data1,
			"Maintinfo": data2,
		},
	})
}
func SoftwareUpdate(c *gin.Context) { //TODO:对于没有填入的项目 如何区分呢
	req := models.Software{}
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("Bad request")
		c.Status(http.StatusBadRequest)
		return
	}
	set := map[string]bool{
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
	query := dbc.DB().Model(&models.Software{}).Where("Software_id = ?", req.SoftwareID)
	queryM := dbc.DB().Model(&models.SoftwareMaintenance{}).Where("Software_id = ? and status = '待处理'", req.SoftwareID)
	chk := new(models.SoftwareMaintenance)
	err := queryM.First(chk).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err,
		})
		return
	} else if req.Status != "非正常" && err == nil {
		c.JSON(http.StatusFailedDependency, gin.H{
			"msg": "Attempting to transform an abnormal Software with maintenance processes unfinished to a normal one.",
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

func SoftwareDelete(c *gin.Context) {
	req := struct {
		SoftwareID uint
	}{}
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("Bad request")
		c.Status(http.StatusBadRequest)
		return
	}
	query := dbc.DB().Model(&models.Software{}).Where("Software_id = ?", req.SoftwareID)
	queryM := dbc.DB().Model(&models.SoftwareMaintenance{}).Where("Software_id = ?", req.SoftwareID)
	var tgt models.Software
	if err := query.Find(&tgt).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err,
		})
		return
	}
	chk := new(models.SoftwareMaintenance)
	err := queryM.First(chk).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err,
		})
		return
	} else if err == nil {
		c.JSON(http.StatusFailedDependency, gin.H{
			"msg": "Attempting to remove a Software with maintenance.",
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

func SoftwareCreate(c *gin.Context) {
	req := models.Software{}
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("Bad request")
		c.Status(http.StatusBadRequest)
		return
	}
	set := map[string]bool{
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
