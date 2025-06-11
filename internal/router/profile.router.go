package router

import (
	"chapapp-backend-api/internal/middleware"
	"chapapp-backend-api/internal/wire"

	"github.com/gin-gonic/gin"
)

type ProfileRouter struct {
}

func (userRouter *ProfileRouter) InitProfileRouter(router *gin.RouterGroup) {
	//public router
	profileController, _ := wire.InitModuleProfile()
	profilePublicRouter := router.Group("/profiles")
	{
		//update
		profilePublicRouter.PUT("/:id", middleware.UploadProfileAccountToCloudinary(), profileController.Update)
	}
	// //private router
	// userPrivateRouter := router.Group("/user")
	// // userPrivateRouter.Use(middleware.Limiter())
	// // userPrivateRouter.Use(middleware.AuthMiddleware())
	// // userPrivateRouter.Use(middleware.PermissionMiddleware())
	// {
	// 	userPrivateRouter.GET("/getInfo/:id")
	// }
}
