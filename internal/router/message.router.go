package router

import (
	"chapapp-backend-api/internal/middleware"
	"chapapp-backend-api/internal/wire"

	"github.com/gin-gonic/gin"
)

type MessageRouter struct {
}

func (messageRouter *MessageRouter) InitMessageRouter(router *gin.RouterGroup) {
	//public router
	messageController, _ := wire.InitModuleMessage()
	messagePublicRouter := router.Group("/messages", middleware.AuthMiddleware())
	{
		//update
		messagePublicRouter.POST("", messageController.Create)
		messagePublicRouter.GET("/me/:id", messageController.GetMessages)
		messagePublicRouter.DELETE("", messageController.Delete)
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
