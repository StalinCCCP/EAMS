package service

import (
	"EAMSbackend/models"
	"EAMSbackend/util"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func HardwareMaintenanceListQuery(c *gin.Context) {
	req := struct {
		HardwareName        string
		HardwareCategory    string
		HardwareStatus      string
		HardwareLocation    util.NullString
		MaintenanceDateFrom time.Time
		MaintenanceDateTo   time.Time
		CostFrom            time.Time
		CostTo              time.Time
		MaintenanceStatus   string
	}{}
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("Bad request")
		c.Status(http.StatusBadRequest)
		return
	}
	data := models.HDLQreq{
		HardwareName: req.HardwareName,
		Category:     req.HardwareCategory,
		Status:       req.HardwareStatus,
		Location:     req.HardwareLocation,
	}
	c.Set("LocalCallData", data)
	HardwareListQuery(c)
	transfer, ok := c.Get("LocalCallResult")
	delete(c.Keys, "LocalCallResult")
	if !ok {
		c.Status(http.StatusInternalServerError)
		return
	}
	buf := transfer.([]models.Hardware)
	var id []uint
	for _, HD := range buf {
		id = append(id, HD.HardwareID)
	}
}
