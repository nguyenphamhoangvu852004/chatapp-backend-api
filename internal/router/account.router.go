package router

import (
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
		authPublicRouter.GET("/", authController.GetList)
		authPublicRouter.GET("/detail/:id", authController.GetDetail)
	}
	{

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
