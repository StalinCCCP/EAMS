package service

import (
	"EAMSbackend/models"
	"EAMSbackend/util"
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

	if err := c.ShouldBindJSON(&DBconf); err != nil {
		log.Println("Bad request")
		c.Status(http.StatusBadRequest)
		return
	}
	DBconf.Userrole = "Supervisor"
	DBconf.Entry_date = time.Now()
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
	// sqlfile, err := os.ReadFile("sql/init.sql")
	// if err != nil {
	// 	log.Println("Failed to open SQL file:", err)
	// 	c.Status(http.StatusInternalServerError)
	// 	os.Remove("DBinfo.json")
	// 	return
	// }
	// sqlstmt := string(sqlfile)
	// err = db.Exec(sqlstmt).Error
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Println("Failed to create tables:", err)
		c.Status(http.StatusInternalServerError)
		os.Remove("DBinfo.json")
		return
	}
	err = db.AutoMigrate(&models.Hardware{})
	if err != nil {
		log.Println("Failed to create tables:", err)
		c.Status(http.StatusInternalServerError)
		os.Remove("DBinfo.json")
		return
	}
	err = db.AutoMigrate(&models.HardwareMaintenance{})
	if err != nil {
		log.Println("Failed to create tables:", err)
		c.Status(http.StatusInternalServerError)
		os.Remove("DBinfo.json")
		return
	}
	err = db.AutoMigrate(&models.Lab{})
	if err != nil {
		log.Println("Failed to create tables:", err)
		c.Status(http.StatusInternalServerError)
		os.Remove("DBinfo.json")
		return
	}
	err = db.AutoMigrate(&models.LabMaintenance{})
	if err != nil {
		log.Println("Failed to create tables:", err)
		c.Status(http.StatusInternalServerError)
		os.Remove("DBinfo.json")
		return
	}
	err = db.AutoMigrate(&models.Software{})
	if err != nil {
		log.Println("Failed to create tables:", err)
		c.Status(http.StatusInternalServerError)
		os.Remove("DBinfo.json")
		return
	}
	err = db.AutoMigrate(&models.SoftwareMaintenance{})
	if err != nil {
		log.Println("Failed to create tables:", err)
		c.Status(http.StatusInternalServerError)
		os.Remove("DBinfo.json")
		return
	}
	supervisor := &models.User{
		Username:     DBconf.Username,
		Pwd:          util.GetMd5(DBconf.Pwd),
		Userrole:     DBconf.Userrole,
		Full_name:    DBconf.Full_name,
		Email:        DBconf.Email,
		Phone_number: DBconf.Phone_number,
		Entry_date:   DBconf.Entry_date,
	}
	err = db.Create(supervisor).Error
	if err != nil {
		log.Println("Failed to create supervisor:", err)
		c.Status(http.StatusInternalServerError)
		os.Remove("DBinfo.json")
		return
	}
	token, err := util.GenerateToken(supervisor.User_id, supervisor.Username, supervisor.Userrole)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Failed to generate token",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"data": map[string]interface{}{
			"token": token,
		}})
}
