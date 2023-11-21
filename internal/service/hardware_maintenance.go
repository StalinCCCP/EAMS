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

func HardwareMaintenanceListQuery(c *gin.Context) {
	req := struct {
		HardwareName        string
		HardwareCategory    string
		HardwareStatus      string
		HardwareLocation    util.NullString
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
	post := define.HLQreq{
		HardwareName: req.HardwareName,
		Category:     req.HardwareCategory,
		Status:       req.HardwareStatus,
		Location:     req.HardwareLocation,
	}
	c.Set("LocalCallData", post)
	HardwareListQuery(c)
	transfer, ok := c.Get("LocalCallResult")
	delete(c.Keys, "LocalCallResult")
	if !ok {
		c.Status(http.StatusInternalServerError)
		return
	}
	buf := transfer.([]define.HLQresp)
	var id []uint
	for _, HD := range buf {
		id = append(id, HD.HardwareID)
	}
	query := dbc.DB().Model(&models.HardwareMaintenance{})
	if id != nil {
		query = query.Where("hardware_id in (?)", id)
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
	var data []define.HMLQresp
	err := query.Select("maintenance_process_id,hardware_id,maintenance_date,cost,status").Scan(&data).Error
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func HardwareMaintenanceUpdate(c *gin.Context) { //TODO:对于没有填入的项目 如何区分呢 另外 需要用反射实现只提交提交过来的项
	req := models.HardwareMaintenance{}
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
	query := dbc.DB().Model(&models.HardwareMaintenance{}).Where("maintenance_process_id = ?", req.MaintenanceProcessID)
	var HardwareID []uint
	err := query.Pluck("hardware_id", &HardwareID).Error
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	// 完成的变成未完成则要求对应的Hardware必须非正常
	query = dbc.DB().Model(&models.HardwareMaintenance{}).Where("maintenance_process_id = ?", req.MaintenanceProcessID)
	var OriginalMaintenanceState []string
	err = query.Pluck("status", &OriginalMaintenanceState).Error
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	var HardwareState []string
	query = dbc.DB().Model(&models.Hardware{}).Where("hardware_id = ?", HardwareID[0])
	err = query.Pluck("status", &HardwareState).Error
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	if OriginalMaintenanceState[0] == "已完成" && req.Status == "待处理" && HardwareState[0] != "非正常" {
		c.JSON(http.StatusFailedDependency, gin.H{
			"msg": "Attempting to transform a finished maintenance process to an unfinished one when corresponding hardware has been in normal state.",
		})
		return
	}
	query = dbc.DB().Model(&models.HardwareMaintenance{}).Where("maintenance_process_id = ?", req.MaintenanceProcessID)
	if err = query.Updates(&req).Error; err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}
func HardwareMaintenanceDetailQuery(c *gin.Context) {
	req := struct {
		MaintenanceProcessID uint
	}{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	var data models.HardwareMaintenance
	if err := dbc.DB().Model(&models.HardwareMaintenance{}).Where("maintenance_process_id = ?", req.MaintenanceProcessID).First(&data).Error; err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
func HardwareMaintenanceDelete(c *gin.Context) {
	req := struct {
		MaintenanceProcessID uint
	}{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	query := dbc.DB().Model(&models.HardwareMaintenance{}).Where("maintenance_process_id = ?", req.MaintenanceProcessID)
	var tgt models.HardwareMaintenance
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

func HardwareMaintenanceCreate(c *gin.Context) {
	req := models.HardwareMaintenance{}
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
	query := dbc.DB().Model(&models.HardwareMaintenance{}).Where("hardware_id = ?", req.HardwareID)
	//不存在的hardware不能存在维修过程 正常的hardware不能存在未完成的维修过程
	chk := new(models.Hardware)
	if err := query.First(chk).Error; err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusFailedDependency, gin.H{
			"msg": "This HardwareID does not exist.",
		})
		return
	} else if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	if chk.Status != "非正常" && req.Status == "待处理" {
		c.JSON(http.StatusFailedDependency, gin.H{
			"msg": "Attempting to add an undone maintenance process to a normal hardware.",
		})
		return
	}
	c.Status(http.StatusOK)
}
