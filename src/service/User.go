package service

import (
	"EAMSbackend/dbc"
	"EAMSbackend/helper"
	"EAMSbackend/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetUserDetail
// @Tags 公共方法
// @Summary 根据用户ID找到用户，显示用户详情
// @Param User_id query string true "用户ID"
// @Success 200 {string} json "{"code":"200","data":""} data中包含了用户ID、用户名、用户密码、用户角色、全名、电子邮箱、电话号码和注册日期"
// @Router /user-detail [get]
func GetUserDetail(c *gin.Context) {
	user_id := c.Query("User_id")
	if user_id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "UserID is not null!",
		})
		return
	}
	data := new(models.User)
	err := dbc.DB.Omit("Pwd").Where("User_id = ?", user_id).Find(&data).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "User not found" + user_id,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

// Login
// @Tags 公共方法
// @Summary 根据用户名和密码登录，并返回token
// @Param Username query string true "用户名"
// @Param Pwd query string true "密码"
// @Success 200 {string} json "{"code":"200","data":""} data中包含了用户token"
// @Router /login [post]
func Login(c *gin.Context) {
	username := c.PostForm("Username")
	pwd := c.PostForm("Pwd")
	if username == "" || pwd == "" {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Username or Password should not be null"})
		return
	}
	pwd = helper.GetMd5(pwd)
	data := new(models.User)
	err := dbc.DB.Where("Username = ? AND Pwd = ?", username, pwd).First(&data).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": "Username or Password is wrong",
			})
			return
		}
	}
	token, err := helper.GenerateToken(data.User_id, data.Username, data.Userrole)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Failed to generate token",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": map[string]interface{}{
			"token":    token,
			"userrole": data.Userrole,
		},
	})
}

// Register
// @Tags 公共方法
// @Summary 输入用户名、密码、电子邮箱、手机号来创建用户（默认创建普通用户，提权在其他地方实现），并返回用户token
// @Param Username query string true "用户名"
// @Param Pwd query string true "密码"
// @Param Full_name query string true "全名"
// @Param Email query string true "电子邮箱"
// @Param Phone_number query string true "手机号"
// @Success 200 {string} json "{"code":"200","data":""} data中包含了用户token"
// @Router /register [post]
func Register(c *gin.Context) {
	username := c.PostForm("Username")
	pwd := c.PostForm("Pwd")
	full_name := c.PostForm("Full_name")
	email := c.PostForm("Email")
	phonenum := c.PostForm("Phone_number")
	userrole := "Normal"
	if username == "" || pwd == "" || full_name == "" || email == "" || phonenum == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Wrong parameters",
		})
		return
	}
	var cnt int64
	err := dbc.DB.Where("Email=?", email).Model(new(models.User)).Count(&cnt).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Error getting user",
		})
		return
	}
	if cnt > 0 {
		c.JSON(http.StatusFailedDependency, gin.H{
			"msg": "该邮箱已被注册",
		})
	}
	data := &models.User{
		Username:     username,
		Pwd:          pwd,
		Userrole:     userrole,
		Email:        email,
		Phone_number: phonenum,
		Entry_date:   time.Now(),
	}
	err = dbc.DB.Create(data).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Failed to create user",
		})
		return
	}
	var user_id uint
	dbc.DB.Where("Email = ?", email).First(&user_id)
	token, err := helper.GenerateToken(user_id, username, userrole)
	c.JSON(http.StatusOK, gin.H{
		"data": map[string]interface{}{
			"token": token,
		}})
}
