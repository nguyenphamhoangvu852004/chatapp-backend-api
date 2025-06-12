package router

type RouterGroup struct {
	AuthRouter
	ProfileRouter
	AccountRouter
	FriendShipRouter
	// BlockRouter
	MessageRouter
}

var RouterGroupApp = new(RouterGroup)
