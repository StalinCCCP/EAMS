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
	authUser.POST("/user-detail", service.GetUserDetail)
	authUser.POST("/user-list", service.GetUserList)
	authUser.POST("/hlq", service.HardwareListQuery)
	authUser.GET("/hcq", service.HardwareCategoryQuery)
	authUser.POST("/hdq", service.HardwareDetailQuery)
	authUser.GET("/hlocq", service.HardwareLocationQuery)
	authUser.POST("/hmlq", service.HardwareMaintenanceListQuery)
	authUser.POST("/hmdq", service.HardwareMaintenanceDetailQuery)
	authUser.POST("/slq", service.SoftwareListQuery)
	// authUser.GET("/scq", service.SoftwareCategoryQuery)
	authUser.POST("/sdq", service.SoftwareDetailQuery)
	authUser.GET("/slocq", service.SoftwareLocationQuery)
	authUser.POST("/smlq", service.SoftwareMaintenanceListQuery)
	authUser.POST("/smdq", service.SoftwareMaintenanceDetailQuery)
	authUser.POST("/llq", service.LabListQuery)
	// authUser.GET("/scq", service.SoftwareCategoryQuery)
	authUser.POST("/ldq", service.LabDetailQuery)
	authUser.GET("/llocq", service.LabLocationQuery)
	authUser.POST("/lmlq", service.LabMaintenanceListQuery)
	authUser.POST("/lmdq", service.LabMaintenanceDetailQuery)
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
