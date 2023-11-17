package middlewares

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Precheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := os.Open("DBinfo.json")
		if err != nil {
			c.Abort()
			log.Println("Not initialized, proceed to init")
			c.JSON(http.StatusFailedDependency, gin.H{
				"msg": "Not initialized",
			})
			return
		}
		c.Next()
		// c.Status(http.StatusOK)
	}
}
