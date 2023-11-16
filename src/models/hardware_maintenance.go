package models

import (
	"EAMSbackend/util"
	"time"
)

type HardwareMaintenance struct {
	MaintenanceProcessID uint            `gorm:"primaryKey;autoIncrement"`
	HardwareID           uint            `gorm:"not null"`
	IssueDescription     util.NullString `gorm:"type:text"`
	SolutionDescription  util.NullString `gorm:"type:text"`
	MaintenanceDate      time.Time       `gorm:"not null"` // 这里的类型可以根据实际情况调整
	Cost                 float64         `gorm:"type:decimal(10,2)"`
	Status               string          `gorm:"type:enum('已完成','待处理')"`
	// Hardware             Hardware        `gorm:"foreignKey:HardwareID"`
}
