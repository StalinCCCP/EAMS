package service

import (
	"EAMSbackend/dbc"
	"EAMSbackend/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init(c *gin.Context) {
	var DBconf models.DBConf
	var Supervisor models.User
	if err := c.ShouldBindJSON(&DBconf); err != nil {
		log.Println("Bad request")
		c.Status(http.StatusBadRequest)
	}
	if err := c.ShouldBindJSON(&Supervisor); err != nil {
		log.Println("Bad request")
		c.Status(http.StatusBadRequest)
	}
	Supervisor.Userrole = "Supervisor"
	Supervisor.Entry_date = time.Now()
	_, err := os.Open("DBinfo.json")
	if err == nil {
		log.Println("File created, not permitted to init")
		c.Status(http.StatusInternalServerError)
		return
	}
	file, err := os.Create("DBinfo.json")
	if err != nil {
		log.Println("Failed to create config file:", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	defer file.Close()
	jsonData, err := json.MarshalIndent(DBconf, "", "    ")
	if err != nil {
		log.Println("Failed to write JSON config into the file")
		c.Status(http.StatusInternalServerError)
		os.Remove("DBinfo.json")
		return
	}
	file.Write(jsonData)
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DBconf.User, DBconf.Password, DBconf.Server, DBconf.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Failed to connect to the database:", err)
		c.Status(http.StatusInternalServerError)
		os.Remove("DBinfo.json")
		return
	}
	sqlfile, err := os.ReadFile("sql/init.sql")
	if err != nil {
		log.Println("Failed to open SQL file:", err)
		c.Status(http.StatusInternalServerError)
		os.Remove("DBinfo.json")
		return
	}
	sqlstmt := string(sqlfile)
	err = db.Raw(sqlstmt).Error
	if err != nil {
		log.Println("Failed to execute sql script:", err)
		c.Status(http.StatusInternalServerError)
		os.Remove("DBinfo.json")
		return
	}
	err = dbc.DB.Create(Supervisor).Error
	if err != nil {
		log.Println("Failed to create supervisor:", err)
		c.Status(http.StatusInternalServerError)
		os.Remove("DBinfo.json")
		return
	}
	c.Status(http.StatusOK)
}
