package manager

import "github.com/gin-gonic/gin"

type AdminRouter struct {
}

func (adminRouter *AdminRouter) InitAdminRouter(router *gin.RouterGroup) {
	// public router
	publicRouter := router.Group("/admin")
	{
		publicRouter.POST("/login")
		// publicRouter.POST("/sendOtp")
	}
	//private router

	privateRouter := router.Group("/admin/user")
	// privateRouter.Use(middleware.Limiter())
	// privateRouter.Use(middleware.AuthMiddleware())
	// privateRouter.Use(middleware.PermissionMiddleware())
	{
		privateRouter.GET("/activeUser/:id")
	}

}
