package router

import (
	"chapapp-backend-api/internal/wire"

	"github.com/gin-gonic/gin"
)

type ConversationRouter struct {
}

func (conversationRouter *ConversationRouter) InitConversationRouter(router *gin.RouterGroup) {
	//public router
	conversationController, _ := wire.InitModuleConversation()
	conversationPublicRouter := router.Group("/conversations")
	{
		//create
		conversationPublicRouter.POST("/", conversationController.Create)
		conversationPublicRouter.POST("/members", conversationController.AddMembers)
		conversationPublicRouter.GET("/owned/me/:id", conversationController.GetListOwnedByMe)
		conversationPublicRouter.GET("/joined/me/:id", conversationController.GetGroupsJoinedByMe)
		conversationPublicRouter.DELETE("/", conversationController.Delete)
		conversationPublicRouter.DELETE("/members", conversationController.DeleteMembers)
		//Status
		// conversationPublicRouter.PUT("/", conversationController.Update)
		// //Get List
		// conversationPublicRouter.GET("/", conversationController.GetList)
		// // Lấy danh sách những người đã trở thành bạn bè với AccountId này
		// conversationPublicRouter.GET("/:id", conversationController.GetListFriendShipsOfAccount)
		// //Lấy danh sách những người mà accountId này đã gữi lời mời tới cho các account khác
		// conversationPublicRouter.GET("/send/:id", conversationController.GetListSendFriendShips)
		// conversationPublicRouter.GET("/receive/:id", conversationController.GetListReceiveFriendShips)
		// conversationPublicRouter.DELETE("/", conversationController.Delete)
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
