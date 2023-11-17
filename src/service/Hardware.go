package service

import (
	"EAMSbackend/dbc"
	"EAMSbackend/models"
	"EAMSbackend/util"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HardwareCategoryQuery(c *gin.Context) {
	var data []util.NullString
	err := dbc.DB().Model(&models.Hardware{}).Distinct("Category").Pluck("Category", &data).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
func HardwareStatusQuery(c *gin.Context) {
	var data []util.NullString
	err := dbc.DB().Model(&models.Hardware{}).Distinct("status").Pluck("status", &data).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
func HardwareLocationQuery(c *gin.Context) {
	var data []util.NullString
	err := dbc.DB().Model(&models.Hardware{}).Distinct("location").Pluck("location", &data).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
func HardwareListQuery(c *gin.Context) {
	name := struct {
		HardwareName string
		Category     string
		Status       string
		Location     util.NullString
	}{}
	if err := c.ShouldBindJSON(&name); err != nil {
		log.Println("Bad request")
		c.Status(http.StatusBadRequest)
		return
	}
	query := dbc.DB().Model(&models.Hardware{}).Where("hardware_name LIKE ?", "%"+name.HardwareName+"%")
	if name.Category != "" {
		query = query.Where("Category = ?", name.Category)
	}
	if name.Status != "" {
		query = query.Where("Status = ?", name.Status)
	}
	if name.Location.String != "" {
		query = query.Where("Location = ?", name.Location)
	}
	var data []models.Hardware
	err := query.Find(&data).Error
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
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
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
	query := dbc.DB().Model(&models.Hardware{}).Where("hardware_id = ?", req)
	err := query.Find(&data1).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err,
		})
		return
	}
	var data2 []models.HardwareMaintenance
	query = dbc.DB().Model(&models.HardwareMaintenance{}).Where("hardware_id = ?", req)
	err = query.Find(&data2).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": map[string]interface{}{
			"HDinfo":    data1,
			"Maintinfo": data2,
		},
	})
}
func HardwareUpdate(c *gin.Context) {
	req := models.Hardware{}
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("Bad request")
		c.Status(http.StatusBadRequest)
		return
	}
	query := dbc.DB().Model(&models.Hardware{}).Where("hardware_id = ?", req.HardwareID)
	queryM := dbc.DB().Model(&models.Hardware{}).Where("hardware_id = ? and status = '待处理'", req.HardwareID)
	var chk *models.HardwareMaintenance
	if err := queryM.First(chk); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err,
		})
		return
	}

	if set := map[string]bool{"保留": true, "正常": true, "占用": true}; set[req.Status] && chk != nil {
		c.JSON(http.StatusFailedDependency, gin.H{
			"msg": "Attempting to transform an abnormal hardware with maintenance processes unfinished to a normal one.",
		})
		return
	}
	if err := query.Updates(&req); err != nil {
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
	queryM := dbc.DB().Model(&models.Hardware{}).Where("hardware_id = ?", req.HardwareID)
	var tgt models.Hardware
	if err := query.Find(&tgt).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err,
		})
		return
	}
	var chk *models.HardwareMaintenance
	if err := queryM.First(chk).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err,
		})
		return
	}
	if chk != nil {
		c.JSON(http.StatusFailedDependency, gin.H{
			"msg": "Attempting to remove a hardware with maintenance unfinished.",
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
	if err := dbc.DB().Create(&req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err,
		})
		return
	}
	c.Status(http.StatusOK)
}
