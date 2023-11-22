package middlewares

import (
	"EAMSbackend/dbc"
	"EAMSbackend/models"
	"EAMSbackend/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthAdminCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		userClaim, err := util.AnalyseToken(auth)
		if err != nil {
			c.Abort()
			c.Status(http.StatusInternalServerError)
			return
		}
		query := dbc.DB().Model(&models.User{}).Where("user_id = ?", userClaim.User_id)
		var UserName []string
		var UserRole []string
		var User_id []int
		if err = query.Pluck("user_id", &User_id).Error; err != nil {
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": "Unauthorized",
			})
		}
		query = dbc.DB().Model(&models.User{}).Where("user_id = ?", userClaim.User_id)

		if err = query.Pluck("username", &UserName).Error; err != nil {
			c.Abort()
			c.Status(http.StatusInternalServerError)
			return
		}
		query = dbc.DB().Model(&models.User{}).Where("user_id = ?", userClaim.User_id)

		if err = query.Pluck("userrole", &UserRole).Error; err != nil {
			c.Abort()
			c.Status(http.StatusInternalServerError)
			return
		}
		if len(UserName) == 0 || len(UserRole) == 0 || len(User_id) == 0 {
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": "Unauthorized",
			})
			return
		}
		if UserName[0] != userClaim.Username || UserRole[0] != userClaim.Userrole {
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": "Unauthorized",
			})
			return
		}
		if userClaim == nil || userClaim.Userrole == "Normal" {
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": "Unauthorized",
			})
			return
		}
		c.Set("Username", UserName[0])
		c.Next()
	}
}

func AuthSupervisorCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		userClaim, err := util.AnalyseToken(auth)
		if err != nil {
			c.Abort()
			c.Status(http.StatusInternalServerError)
			return
		}
		query := dbc.DB().Model(&models.User{}).Where("user_id = ?", userClaim.User_id)
		var UserName []string
		var UserRole []string
		var User_id []int
		if err = query.Pluck("user_id", &User_id).Error; err != nil {
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": "Unauthorized",
			})
		}
		query = dbc.DB().Model(&models.User{}).Where("user_id = ?", userClaim.User_id)

		if err = query.Pluck("username", &UserName).Error; err != nil {
			c.Abort()
			c.Status(http.StatusInternalServerError)
			return
		}
		query = dbc.DB().Model(&models.User{}).Where("user_id = ?", userClaim.User_id)

		if err = query.Pluck("userrole", &UserRole).Error; err != nil {
			c.Abort()
			c.Status(http.StatusInternalServerError)
			return
		}
		if len(UserName) == 0 || len(UserRole) == 0 || len(User_id) == 0 {
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": "Unauthorized",
			})
			return
		}
		if UserName[0] != userClaim.Username || UserRole[0] != userClaim.Userrole {
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": "Unauthorized",
			})
			return
		}
		if userClaim == nil || userClaim.Userrole != "Supervisor" {
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": "Unauthorized",
			})
			return
		}
		c.Next()
	}
}

func AuthUserCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		userClaim, err := util.AnalyseToken(auth)
		if err != nil {
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": "Unauthorized",
			})
			return
		}
		query := dbc.DB().Model(&models.User{}).Where("user_id = ?", userClaim.User_id)
		var UserName []string
		var UserRole []string
		var User_id []int
		if err = query.Pluck("user_id", &User_id).Error; err != nil {
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": "Unauthorized",
			})
			return
		}
		query = dbc.DB().Model(&models.User{}).Where("user_id = ?", userClaim.User_id)
		if err = query.Pluck("username", &UserName).Error; err != nil {
			c.Abort()
			c.Status(http.StatusInternalServerError)
			return
		}
		query = dbc.DB().Model(&models.User{}).Where("user_id = ?", userClaim.User_id)
		if err = query.Pluck("userrole", &UserRole).Error; err != nil {
			c.Abort()
			c.Status(http.StatusInternalServerError)
			return
		}
		if len(UserName) == 0 || len(UserRole) == 0 || len(User_id) == 0 {
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": "Unauthorized",
			})
			return
		}
		if UserName[0] != userClaim.Username || UserRole[0] != userClaim.Userrole {
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": "Unauthorized",
			})
			return
		}
		if userClaim == nil {
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": "Unauthorized",
			})
			return
		}
		c.Next()
	}
}
