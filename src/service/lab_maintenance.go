package service

import (
	"EAMSbackend/dbc"
	"EAMSbackend/define"
	"EAMSbackend/models"
	"EAMSbackend/util"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func LabMaintenanceListQuery(c *gin.Context) {
	req := struct {
		LabName             string
		LabVersion          string
		LabStatus           string
		LabLocation         util.NullString
		MaintenanceDateFrom time.Time
		MaintenanceDateTo   time.Time
		CostFrom            float64
		CostTo              float64
		MaintenanceStatus   string
	}{
		CostFrom: define.Float64Null,
		CostTo:   define.Float64Null,
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("Bad request")
		c.Status(http.StatusBadRequest)
		return
	}
	post := define.LLQreq{
		LabName: req.LabName,
		// Version:  req.LabVersion,
		Status:   req.LabStatus,
		Location: req.LabLocation,
	}
	c.Set("LocalCallData", post)
	LabListQuery(c)
	transfer, ok := c.Get("LocalCallResult")
	delete(c.Keys, "LocalCallResult")
	if !ok {
		c.Status(http.StatusInternalServerError)
		return
	}
	buf := transfer.([]define.LLQresp)
	var id []uint
	for _, HD := range buf {
		id = append(id, HD.LabID)
	}
	query := dbc.DB().Model(&models.LabMaintenance{})
	if id != nil {
		query = query.Where("Lab_id in (?)", id)
	}
	if req.MaintenanceDateFrom != define.TimeNull {
		query.Where("maintenance_date >= ?", req.MaintenanceDateFrom)
	}
	if req.MaintenanceDateTo != define.TimeNull {
		query.Where("maintenance_date <= ?", req.MaintenanceDateTo)
	}
	if req.CostFrom != define.Float64Null {
		query.Where("cost >= ?", req.CostFrom)
	}
	if req.CostTo != define.Float64Null {
		query.Where("cost <= ?", req.CostTo)
	}
	if req.MaintenanceStatus != "" {
		query.Where("status = ?", req.MaintenanceStatus)
	}
	var data []define.LMLQresp
	err := query.Select("maintenance_process_id,Lab_id,maintenance_date,cost,status").Scan(&data).Error
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func LabMaintenanceUpdate(c *gin.Context) { //TODO:对于没有填入的项目 如何区分呢 另外 需要用反射实现只提交提交过来的项
	req := models.LabMaintenance{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	set := map[string]bool{
		"已完成": true,
		"待处理": true,
	}
	if !set[req.Status] {
		log.Println("Bad request")
		c.Status(http.StatusBadRequest)
		return
	}
	query := dbc.DB().Model(&models.LabMaintenance{}).Where("maintenance_process_id = ?", req.MaintenanceProcessID)
	var LabID []uint
	err := query.Pluck("Lab_id", &LabID).Error
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	// 完成的变成未完成则要求对应的Lab必须非正常
	query = dbc.DB().Model(&models.LabMaintenance{}).Where("maintenance_process_id = ?", req.MaintenanceProcessID)
	var OriginalMaintenanceState []string
	err = query.Pluck("status", &OriginalMaintenanceState).Error
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	var LabState []string
	query = dbc.DB().Model(&models.Lab{}).Where("Lab_id = ?", LabID[0])
	err = query.Pluck("status", &LabState).Error
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	if OriginalMaintenanceState[0] == "已完成" && req.Status == "待处理" && LabState[0] != "停用" {
		c.JSON(http.StatusFailedDependency, gin.H{
			"msg": "Attempting to transform a finished maintenance process to an unfinished one when corresponding Lab has been in normal state.",
		})
		return
	}
	query = dbc.DB().Model(&models.LabMaintenance{}).Where("maintenance_process_id = ?", req.MaintenanceProcessID)
	if err = query.Updates(&req).Error; err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}

func LabMaintenanceDelete(c *gin.Context) {
	req := struct {
		MaintenanceProcessID uint
	}{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	query := dbc.DB().Model(&models.LabMaintenance{}).Where("maintenance_process_id = ?", req.MaintenanceProcessID)
	var tgt models.LabMaintenance
	if err := query.Find(&tgt).Error; err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	if err := dbc.DB().Delete(&tgt).Error; err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}

func LabMaintenanceCreate(c *gin.Context) {
	req := models.LabMaintenance{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	set := map[string]bool{
		"已完成": true,
		"待处理": true,
	}
	if !set[req.Status] {
		log.Println("Bad request")
		c.Status(http.StatusBadRequest)
		return
	}
	query := dbc.DB().Model(&models.LabMaintenance{}).Where("Lab_id = ?", req.LabID)
	//不存在的Lab不能存在维修过程 正常的Lab不能存在未完成的维修过程
	chk := new(models.Lab)
	if err := query.First(chk).Error; err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusFailedDependency, gin.H{
			"msg": "This LabID does not exist.",
		})
		return
	} else if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	if chk.Status != "停用" && req.Status == "待处理" {
		c.JSON(http.StatusFailedDependency, gin.H{
			"msg": "Attempting to add an undone maintenance process to a normal Lab.",
		})
		return
	}
	c.Status(http.StatusOK)
}
