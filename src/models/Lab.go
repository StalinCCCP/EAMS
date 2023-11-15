package models

type Lab struct {
	LabID       uint   `gorm:"primaryKey;autoIncrement"`
	LabName     string `gorm:"not null"`
	Description string `gorm:"type:text"`
	Status      string `gorm:"type:enum('正常','停用','占用')"`
}
