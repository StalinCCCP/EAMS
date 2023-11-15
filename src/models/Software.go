package models

type Software struct {
	SoftwareID   uint   `gorm:"primaryKey;autoIncrement"`
	SoftwareName string `gorm:"not null"`
	Version      string
	Description  string `gorm:"type:text"`
	Status       string `gorm:"type:enum('正常','非正常')"`
	Location     string
}
