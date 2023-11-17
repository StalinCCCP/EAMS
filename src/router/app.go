package router

import (
	_ "EAMSbackend/docs"
	"EAMSbackend/middlewares"
	"EAMSbackend/service"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/precheck", service.Precheck)
	r.POST("/init", service.Init)
	r.GET("/user-detail", service.GetUserDetail)
	r.POST("/login", service.Login)
	r.POST("/register", service.Register)
	authUser := r.Group("/user", middlewares.AuthUserCheck())
	authUser.GET("/hdq", service.HardwareListQuery)
	authUser.GET("/hdcq", service.HardwareCategoryQuery)

	return r
}
