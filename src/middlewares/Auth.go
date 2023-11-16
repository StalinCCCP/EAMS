package middlewares

import (
	"EAMSbackend/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthAdminCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		userClaim, err := helper.AnalyseToken(auth)
		if err != nil {
			c.Abort()
			c.Status(http.StatusInternalServerError)
			return
		}
		if userClaim == nil || userClaim.Userrole == "Normal" {
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": "Unauthorized",
			})
			return
		}
		c.Next()
	}
}

func AuthSupervisorCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		userClaim, err := helper.AnalyseToken(auth)
		if err != nil {
			c.Abort()
			c.Status(http.StatusInternalServerError)
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
		userClaim, err := helper.AnalyseToken(auth)
		if err != nil {
			c.Abort()
			c.Status(http.StatusInternalServerError)
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
