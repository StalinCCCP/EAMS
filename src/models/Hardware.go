package models

type Hardware struct {
	HardwareID   uint   `gorm:"primaryKey;autoIncrement"`
	HardwareName string `gorm:"not null"`
	Category     string
	Description  string `gorm:"type:text"`
	Status       string `gorm:"type:enum('保留','正常','占用','非正常')"`
	Location     string
}
