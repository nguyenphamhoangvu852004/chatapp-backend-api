package manager

import (
	"chapapp-backend-api/internal/wire"

	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (userRouter *UserRouter) InitUserRouter(router *gin.RouterGroup) {
	//public router
	userController, _ := wire.InitModuleUser()
	publicRouter := router.Group("/admin/user")
	{
		publicRouter.POST("/register", userController.Register)
		publicRouter.POST("/sendOtp")
	}
	//private router

	privateRouter := router.Group("/admin/user")
	// privateRouter.Use(middleware.Limiter())
	// privateRouter.Use(middleware.AuthMiddleware())
	// privateRouter.Use(middleware.PermissionMiddleware())
	{
		privateRouter.POST("/activeUser/:id")
	}

}
