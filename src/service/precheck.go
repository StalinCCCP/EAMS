package service

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Precheck(c *gin.Context) {
	_, err := os.Open("DBinfo.json")
	if err != nil {
		log.Println("Not initialized, proceed to init")
		c.Status(http.StatusFailedDependency)
	} else {
		c.Status(http.StatusOK)
	}
}
