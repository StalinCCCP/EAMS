package router

import (
	_ "EAMSbackend/docs"
	"EAMSbackend/middlewares"
	"EAMSbackend/service"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	R := gin.Default()
	R.POST("/init", service.Init)
	r := R.Group("/p", middlewares.Precheck())
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/user-detail", service.GetUserDetail)
	r.POST("/login", service.Login)
	r.POST("/register", service.Register)
	authUser := r.Group("/user", middlewares.AuthUserCheck())
	authUser.GET("/hdlq", service.HardwareListQuery)
	authUser.GET("/hdcq", service.HardwareCategoryQuery)
	authUser.GET("/hddq", service.HardwareDetailQuery)
	authUser.GET("/hdlocq", service.HardwareLocationQuery)
	authUser.GET("/hdsq", service.HardwareStatusQuery)
	authUser.POST("/hdupd", service.HardwareUpdate)
	authUser.DELETE("/hddlt", service.HardwareDelete)
	authUser.PUT("/hdc", service.HardwareCreate)
	return R
}
