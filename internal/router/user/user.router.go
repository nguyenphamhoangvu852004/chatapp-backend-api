package user

import (
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (userRouter *UserRouter) InitUserRouter(router *gin.RouterGroup) {
	//public router
	userPublicRouter := router.Group("/user")
	{
		userPublicRouter.POST("/register")
		userPublicRouter.POST("/sendOtp")
	}
	//private router

	userPrivateRouter := router.Group("/user")
	// userPrivateRouter.Use(middleware.Limiter())
	// userPrivateRouter.Use(middleware.AuthMiddleware())
	// userPrivateRouter.Use(middleware.PermissionMiddleware())
	{
		userPrivateRouter.GET("/getInfo/:id")
	}

}
