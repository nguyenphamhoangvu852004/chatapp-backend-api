package router

type RouterGroup struct {
	AuthRouter
	ProfileRouter
	AccountRouter
	FriendShipRouter
}

var RouterGroupApp = new(RouterGroup)
