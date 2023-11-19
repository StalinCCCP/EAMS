package main

import (
	"EAMSbackend/router"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	//gin.DefaultWriter = io.Discard
	r := router.Router()
	r.Run()
}
