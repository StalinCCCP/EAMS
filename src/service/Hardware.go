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
func HardwareQuery(c *gin.Context) {
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
	query := dbc.DB().Where("hardware_name LIKE ?", "%"+name.HardwareName+"%")
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
	if data == nil {
		c.JSON(http.StatusOK, gin.H{
			"data": "",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	}
}
