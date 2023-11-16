package models

import "time"

type DBConf struct {
	Server       string `json:"Server"`
	DBName       string `json:"DBName"`
	User         string `json:"User"`
	Password     string `json:"Password"`
	Username     string
	Pwd          string
	Userrole     string `gorm:"type:enum(enum('Normal','Admin','Supervisor')"`
	Full_name    string
	Email        string
	Phone_number string
	Entry_date   time.Time
}
