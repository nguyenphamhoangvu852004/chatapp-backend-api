package router

import (
	"chapapp-backend-api/internal/middleware"
	"chapapp-backend-api/internal/wire"

	"github.com/gin-gonic/gin"
)

type AccountRouter struct {
}

func (accountRouter *AccountRouter) InitAccountRouter(router *gin.RouterGroup) {
	//public router
	authController, _ := wire.InitModuleAccount()
	authPublicRouter := router.Group("/accounts")
	{
		authPublicRouter.GET("", authController.GetList)
		authPublicRouter.GET("/detail/:id", authController.GetDetail)
		authPublicRouter.GET("/random", authController.GetRandomList)
	}
	{
		authPublicRouter.PATCH("", middleware.AuthMiddleware(), authController.ChangePassword)
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
