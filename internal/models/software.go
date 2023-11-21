package models

import (
	"EAMSbackend/util"
)

type Software struct {
	SoftwareID   uint   `gorm:"primaryKey;autoIncrement"`
	SoftwareName string `gorm:"not null"`
	Version      util.NullString
	Description  util.NullString `gorm:"type:text"`
	Status       string          `gorm:"type:enum('正常','非正常')"`
	Location     util.NullString
}
