package router

import (
	"chapapp-backend-api/internal/middleware"
	"chapapp-backend-api/internal/wire"

	"github.com/gin-gonic/gin"
)

type BlockRouter struct {
}

func (blockRouter *BlockRouter) InitBlockRouter(router *gin.RouterGroup) {
	//public router
	blockController, _ := wire.InitModuleBlock()
	blockPublicRouter := router.Group("/blocks", middleware.AuthMiddleware())
	{
		blockPublicRouter.POST("/", blockController.Create)
		blockPublicRouter.DELETE("/", blockController.Delete)
		blockPublicRouter.GET("/me/:id", blockController.GetList)
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
