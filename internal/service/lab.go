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

// func LabVersionQuery(c *gin.Context) {
// 	var data []util.NullString
// 	err := dbc.DB().Model(&models.Lab{}).Distinct("Version").Pluck("Version", &data).Error
// 	if err != nil {
// 		c.Status(http.StatusInternalServerError)
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"data": data,
// 	})
// }

//	func LabStatusQuery(c *gin.Context) {
//		var data []util.NullString
//		err := dbc.DB().Model(&models.Lab{}).Distinct("status").Pluck("status", &data).Error
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
func LabLocationQuery(c *gin.Context) {
	var data []util.NullString
	err := dbc.DB().Model(&models.Lab{}).Distinct("location").Pluck("location", &data).Error
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
func LabListQuery(c *gin.Context) {
	name := define.LLQreq{}
	transfer, ok := c.Get("LocalCallData")
	if !ok {
		if err := c.ShouldBindJSON(&name); err != nil {
			log.Println("Bad request")
			c.Status(http.StatusBadRequest)
			return
		}
	} else {
		name = transfer.(define.LLQreq)
		delete(c.Keys, "LocalCallData")
	}
	query := dbc.DB().Model(&models.Lab{})
	if name.LabName != "" {
		query = query.Where("Lab_name LIKE ?", "%"+name.LabName+"%")
	}
	// if name.Version != "" {
	// 	query = query.Where("Version = ?", name.Version)
	// }
	if name.Status != "" {
		query = query.Where("Status = ?", name.Status)
	}
	if name.Location.String != "" {
		query = query.Where("Location = ?", name.Location)
	}
	var data []define.LLQresp
	err := query.Select("Lab_id, Lab_name,status,location").Scan(&data).Error
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

func LabDetailQuery(c *gin.Context) {
	req := struct {
		LabID uint
	}{}
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("Bad request")
		c.Status(http.StatusBadRequest)
		return
	}
	var data1 []models.Lab
	query := dbc.DB().Model(&models.Lab{}).Where("Lab_id = ?", req.LabID)
	err := query.Find(&data1).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err,
		})
		return
	}
	var data2 models.LabMaintenance
	query = dbc.DB().Model(&models.LabMaintenance{}).Where("Lab_id = ?", req.LabID)
	err = query.First(&data2).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": map[string]interface{}{
			"Linfo":     data1,
			"Maintinfo": data2,
		},
	})
}
func LabUpdate(c *gin.Context) { //TODO:对于没有填入的项目 如何区分呢
	req := models.Lab{}
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("Bad request")
		c.Status(http.StatusBadRequest)
		return
	}
	set := map[string]bool{
		"占用": true,
		"正常": true,
		"停用": true,
	}
	if !set[req.Status] {
		log.Println("Bad request")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Status is not defined",
		})
		return
	}
	query := dbc.DB().Model(&models.Lab{}).Where("Lab_id = ?", req.LabID)
	queryM := dbc.DB().Model(&models.LabMaintenance{}).Where("Lab_id = ? and status = '待处理'", req.LabID)
	chk := new(models.LabMaintenance)
	err := queryM.First(chk).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err,
		})
		return
	} else if req.Status != "停用" && err == nil {
		c.JSON(http.StatusFailedDependency, gin.H{
			"msg": "Attempting to transform an abnormal Lab with maintenance processes unfinished to a normal one.",
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

func LabDelete(c *gin.Context) {
	req := struct {
		LabID uint
	}{}
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("Bad request")
		c.Status(http.StatusBadRequest)
		return
	}
	query := dbc.DB().Model(&models.Lab{}).Where("Lab_id = ?", req.LabID)
	queryM := dbc.DB().Model(&models.LabMaintenance{}).Where("Lab_id = ?", req.LabID)
	var tgt models.Lab
	if err := query.Find(&tgt).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err,
		})
		return
	}
	chk := new(models.LabMaintenance)
	err := queryM.First(chk).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err,
		})
		return
	} else if err == nil {
		c.JSON(http.StatusFailedDependency, gin.H{
			"msg": "Attempting to remove a Lab with maintenance.",
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

func LabCreate(c *gin.Context) {
	req := models.Lab{}
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("Bad request")
		c.Status(http.StatusBadRequest)
		return
	}
	set := map[string]bool{
		"占用": true,
		"正常": true,
		"停用": true,
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
