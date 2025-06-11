package router

import (
	"chapapp-backend-api/internal/wire"

	"github.com/gin-gonic/gin"
)

type FriendShipRouter struct {
}

func (friendShipRouter *FriendShipRouter) InitFriendShipRouter(router *gin.RouterGroup) {
	//public router
	friendShipController, _ := wire.InitModuleFriendShip()
	friendShipPublicRouter := router.Group("/friendShips")
	{
		//create
		friendShipPublicRouter.POST("/", friendShipController.Create)
		//Status
		friendShipPublicRouter.PUT("/", friendShipController.Update)
		//Get List
		friendShipPublicRouter.GET("/", friendShipController.GetList)
		// Lấy danh sách những người đã trở thành bạn bè với AccountId này
		friendShipPublicRouter.GET("/:id", friendShipController.GetListFriendShipsOfAccount)
		//Lấy danh sách những người mà accountId này đã gữi lời mời tới cho các account khác
		friendShipPublicRouter.GET("/send/:id", friendShipController.GetListSendFriendShips)
		friendShipPublicRouter.GET("/receive/:id", friendShipController.GetListReceiveFriendShips)

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
