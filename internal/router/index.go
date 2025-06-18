package router

type RouterGroup struct {
	AuthRouter
	ProfileRouter
	AccountRouter
	FriendShipRouter
	BlockRouter
	MessageRouter
	ConversationRouter
	BanRouter
}

var RouterGroupApp = new(RouterGroup)
