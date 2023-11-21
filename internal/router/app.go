package router

import (
	_ "EAMSbackend/docs"
	"EAMSbackend/middlewares"
	"EAMSbackend/service"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	R := gin.Default()
	R.Use(middlewares.Cors())
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
	authUser.GET("/hmdq", service.HardwareMaintenanceDetailQuery)
	authUser.GET("/slq", service.SoftwareListQuery)
	// authUser.GET("/scq", service.SoftwareCategoryQuery)
	authUser.GET("/sdq", service.SoftwareDetailQuery)
	authUser.GET("/slocq", service.SoftwareLocationQuery)
	authUser.GET("/smlq", service.SoftwareMaintenanceListQuery)
	authUser.GET("/smdq", service.SoftwareMaintenanceDetailQuery)
	authUser.GET("/llq", service.LabListQuery)
	// authUser.GET("/scq", service.SoftwareCategoryQuery)
	authUser.GET("/ldq", service.LabDetailQuery)
	authUser.GET("/llocq", service.LabLocationQuery)
	authUser.GET("/lmlq", service.LabMaintenanceListQuery)
	authUser.GET("/lmdq", service.LabMaintenanceDetailQuery)
	// authUser.GET("/hdsq", service.HardwareStatusQuery)
	authAdmin := r.Group("/admin", middlewares.AuthAdminCheck())
	authAdmin.POST("/hupd", service.HardwareUpdate)
	authAdmin.DELETE("/hdlt", service.HardwareDelete)
	authAdmin.PUT("/hc", service.HardwareCreate)
	authAdmin.POST("/hmupd", service.HardwareMaintenanceUpdate)
	authAdmin.PUT("/hmc", service.HardwareMaintenanceCreate)
	authAdmin.DELETE("/hmdlt", service.HardwareMaintenanceDelete)
	authAdmin.POST("/supd", service.SoftwareUpdate)
	authAdmin.DELETE("/sdlt", service.SoftwareDelete)
	authAdmin.PUT("/sc", service.SoftwareCreate)
	authAdmin.POST("/smupd", service.SoftwareMaintenanceUpdate)
	authAdmin.PUT("/smc", service.SoftwareMaintenanceCreate)
	authAdmin.DELETE("/smdlt", service.SoftwareMaintenanceDelete)
	authAdmin.POST("/lupd", service.LabUpdate)
	authAdmin.DELETE("/ldlt", service.LabDelete)
	authAdmin.PUT("/lc", service.LabCreate)
	authAdmin.POST("/lmupd", service.LabMaintenanceUpdate)
	authAdmin.PUT("/lmc", service.LabMaintenanceCreate)
	authAdmin.DELETE("/lmdlt", service.LabMaintenanceDelete)
	authSupervisor := r.Group("/supervisor", middlewares.AuthSupervisorCheck())
	authSupervisor.POST("/chmod", service.ChangeUserRole)
	return R
}
