package models

import (
	"EAMSbackend/util"
	"time"
)

type User struct {
	User_id      uint `gorm:"primaryKey;autoIncrement"`
	Username     string
	Pwd          string
	Userrole     string `gorm:"type:enum('Normal','Admin','Supervisor')"`
	Full_name    util.NullString
	Email        util.NullString
	Phone_number util.NullString
	Entry_date   time.Time
}
