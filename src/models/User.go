package models

import "time"

type User struct {
	User_id      uint `gorm:"primaryKey;autoIncrement"`
	Username     string
	Pwd          string
	Userrole     string
	Full_name    string
	Email        string
	Phone_number string
	Entry_date   time.Time
}
