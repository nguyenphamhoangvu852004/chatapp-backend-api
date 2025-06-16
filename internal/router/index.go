package router

type RouterGroup struct {
	AuthRouter
	ProfileRouter
	AccountRouter
	FriendShipRouter
	// BlockRouter
	MessageRouter
	ConversationRouter
}

var RouterGroupApp = new(RouterGroup)
