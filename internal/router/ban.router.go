package router

import (
	"chapapp-backend-api/internal/middleware"
	"chapapp-backend-api/internal/wire"

	"github.com/gin-gonic/gin"
)

type BanRouter struct {
}

func (banRouter *BanRouter) InitBanRouter(router *gin.RouterGroup) {
	//public router
	banController, _ := wire.InitModuleBan()
	banPublicRouter := router.Group("/bans", middleware.AuthMiddleware())

	{
		banPublicRouter.GET("", banController.GetList)
		banPublicRouter.POST("/create", middleware.VerifyRole([]string{"ADMIN"}), banController.Create)
		banPublicRouter.POST("/delete", middleware.VerifyRole([]string{"ADMIN"}), banController.Delete)
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
