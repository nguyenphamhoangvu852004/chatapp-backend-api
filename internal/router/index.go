package router

type RouterGroup struct {
	AuthRouter
	ProfileRouter
	AccountRouter
}

var RouterGroupApp = new(RouterGroup)
