package dbc

import (
	"EAMSbackend/models"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var c *gorm.DB

func DB() *gorm.DB {
	if c != nil {
		return c
	}
	file, err := os.Open("DBinfo.json")
	if err != nil {
		log.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		log.Println("Error reading file:", err)
		return nil
	}
	var DBconf models.DBConf
	err = json.Unmarshal(data, &DBconf)
	if err != nil {
		log.Println("Error unmarshalling JSON:", err)
		return nil
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DBconf.User, DBconf.Password, DBconf.Server, DBconf.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Error Connecting MySQL:", err)
	}
	c = db
	return db
}
