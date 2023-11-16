package models

import (
	"EAMSbackend/util"
)

type Hardware struct {
	HardwareID   uint            `gorm:"primaryKey;autoIncrement"`
	HardwareName string          `gorm:"not null"`
	Category     string          `gorm:"not null"`
	Description  util.NullString `gorm:"type:text"`
	Status       string          `gorm:"type:enum('保留','正常','占用','非正常')"`
	Location     util.NullString
}
