package models

import (
	"time"
)

type SoftwareMaintenance struct {
	MaintenanceProcessID uint      `gorm:"primaryKey;autoIncrement"`
	SoftwareID           uint      // 注意：这里的 SoftwareID 类型可能需要调整，以匹配 Software 模型的主键类型
	IssueDescription     string    `gorm:"type:text"`
	SolutionDescription  string    `gorm:"type:text"`
	MaintenanceDate      time.Time // 这里的类型可以根据实际情况调整
	Cost                 float64   `gorm:"type:decimal(10,2)"`
	Status               string    `gorm:"type:enum('已完成','待处理')"`
	Software             Software  `gorm:"foreignKey:SoftwareID"`
}
