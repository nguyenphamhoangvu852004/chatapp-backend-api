package router

import (
	"chapapp-backend-api/internal/wire"

	"github.com/gin-gonic/gin"
)

type BanRouter struct {
}

func (banRouter *BanRouter) InitBanRouter(router *gin.RouterGroup) {
	//public router
	banController, _ := wire.InitModuleBan()
	banPublicRouter := router.Group("/bans")

	{
		banPublicRouter.GET("", banController.GetList)
		banPublicRouter.POST("/create", banController.Create)
		banPublicRouter.POST("/delete", banController.Delete)
	}

	// //private router
	// banPrivateRouter := router.Group("/user")
	// // banPrivateRouter.Use(middleware.Limiter())
	// // banPrivateRouter.Use(middleware.BanMiddleware())
	// // banPrivateRouter.Use(middleware.PermissionMiddleware())
	// {
	// 	banPrivateRouter.GET("/getInfo/:id")
	// }

}
