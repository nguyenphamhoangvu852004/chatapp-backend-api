package router

import (
	"chapapp-backend-api/internal/wire"

	"github.com/gin-gonic/gin"
)

type BlockRouter struct {
}

func (blockRouter *BlockRouter) InitBlockRouter(router *gin.RouterGroup) {
	//public router
	blockController, _ := wire.InitModuleBlock()
	blockPublicRouter := router.Group("/blocks")
	{
		blockPublicRouter.POST("/", blockController.Create)
		blockPublicRouter.DELETE("/", blockController.Delete)
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
