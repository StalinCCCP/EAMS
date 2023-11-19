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
	r.POST("/login", service.Login)
	r.POST("/register", service.Register)
	authUser := r.Group("/user", middlewares.AuthUserCheck())
	authUser.GET("/user-detail", service.GetUserDetail)
	authUser.GET("/user-list", service.GetUserList)
	authUser.GET("/hlq", service.HardwareListQuery)
	authUser.GET("/hcq", service.HardwareCategoryQuery)
	authUser.GET("/hdq", service.HardwareDetailQuery)
	authUser.GET("/hlocq", service.HardwareLocationQuery)
	authUser.GET("/hmlq", service.HardwareMaintenanceListQuery)
	// authUser.GET("/hdsq", service.HardwareStatusQuery)
	authAdmin := r.Group("/admin", middlewares.AuthAdminCheck())
	authAdmin.POST("/hupd", service.HardwareUpdate)
	authAdmin.DELETE("/hdlt", service.HardwareDelete)
	authAdmin.PUT("/hc", service.HardwareCreate)
	authAdmin.POST("/hmupd", service.HardwareMaintenanceUpdate)
	authAdmin.PUT("/hmc", service.HardwareMaintenanceCreate)
	authAdmin.DELETE("/hmdlt", service.HardwareMaintenanceDelete)
	authSupervisor := r.Group("/supervisor", middlewares.AuthSupervisorCheck())
	authSupervisor.POST("/chmod", service.ChangeUserRole)
	return R
}
