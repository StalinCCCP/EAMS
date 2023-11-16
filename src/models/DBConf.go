package models

import (
	"EAMSbackend/util"
	"time"
)

type DBConf struct {
	Server       string `json:"Server"`
	DBName       string `json:"DBName"`
	User         string `json:"User"`
	Password     string `json:"Password"`
	Username     string
	Pwd          string
	Userrole     string `gorm:"type:enum(enum('Normal','Admin','Supervisor')"`
	Full_name    util.NullString
	Email        util.NullString
	Phone_number util.NullString
	Entry_date   time.Time
}
