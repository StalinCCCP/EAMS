package service

import (
	"EAMSbackend/dbc"
	"EAMSbackend/define"
	"EAMSbackend/models"
	"EAMSbackend/util"
	"fmt"
	"log"
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
	uid := struct {
		User_id uint
	}{}
	if err := c.ShouldBindJSON(&uid); err != nil {
		log.Println("Bad request")
		c.Status(http.StatusBadRequest)
		return
	}
	//user_id := c.Query("User_id")
	// if user_id == "" {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"msg": "UserID is not null!",
	// 	})
	// 	return
	// }
	data := new(models.User)
	err := dbc.DB().Omit("pwd").Where("user_id = ?", uid.User_id).First(&data).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "User not found " + fmt.Sprintf("%d", uid.User_id),
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
	// username := c.PostForm("Username")
	// pwd := c.PostForm("Pwd")
	user := struct {
		Username string
		Pwd      string
	}{}
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Println("Bad request")
		c.Status(http.StatusBadRequest)
		return
	}
	if user.Username == "" || user.Pwd == "" {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Username or Password should not be null"})
		return
	}
	user.Pwd = util.GetMd5(user.Pwd)
	data := new(models.User)
	err := dbc.DB().Where("Username = ? AND Pwd = ?", user.Username, user.Pwd).First(&data).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": "Username or Password is wrong",
			})
			return
		}
	}
	token, err := util.GenerateToken(data.User_id, data.Username, data.Userrole)
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
	data := new(models.User)
	if err := c.ShouldBindJSON(data); err != nil {
		log.Println("Bad request")
		c.Status(http.StatusBadRequest)
		return
	}
	if data.Username == "" || data.Pwd == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Wrong parameters",
		})
		return
	}
	var cnt int64
	err := dbc.DB().Where("Username=?", data.Username).Model(new(models.User)).Count(&cnt).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Error getting user",
		})
		return
	}
	if cnt > 0 {
		c.JSON(http.StatusFailedDependency, gin.H{
			"msg": "This username has been registered",
		})
		return
	}
	// data := &models.User{
	// 	Username:     username,
	// 	Pwd:          pwd,
	// 	Userrole:     userrole,
	// 	Email:        email,
	// 	Phone_number: phonenum,
	// 	Entry_date:   time.Now(),
	// }
	data.Entry_date = time.Now()
	data.Userrole = "Normal"
	err = dbc.DB().Create(data).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Failed to create user",
		})
		return
	}
	var user_id uint
	dbc.DB().Where("Username = ?", data.Username).First(&user_id)
	token, err := util.GenerateToken(user_id, data.Username, data.Userrole)
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

func GetUserList(c *gin.Context) {
	name := struct {
		Username string
	}{}

	if err := c.ShouldBindJSON(&name); err != nil {
		log.Println("Bad request")
		c.Status(http.StatusBadRequest)
		return
	}

	query := dbc.DB().Model(&models.User{})
	if name.Username != "" {
		query = query.Where("username like ?", "%"+name.Username+"%")
	}
	var data []define.Userresp
	err := query.Select("user_id, username, userrole").Scan(&data).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err,
		})
		return
	}
	// jsonData, err := json.Marshal(data)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"msg": err,
	// 	})
	// 	return
	// }
	//fmt.Println(string(jsonData))
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})

}

func ChangeUserRole(c *gin.Context) {
	req := struct {
		User_id  uint
		Userrole string
	}{}
	set := map[string]bool{
		"Normal": true,
		"Admin":  true,
	}
	if !set[req.Userrole] {
		c.Status(http.StatusBadRequest)
		return
	}
	query := dbc.DB().Model(&models.User{}).Where("user_id = ?", req.User_id)
	if err := query.Update("userrole", req.Userrole).Error; err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}
