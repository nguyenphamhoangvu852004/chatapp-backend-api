package router

import (
	"chapapp-backend-api/internal/router/manager"
	"chapapp-backend-api/internal/router/user"
)

type RouterGroup struct {
	UserRouter    user.UserRouterGroup
	ManagerRouter manager.ManagerRouterGroup
}

var RouterGroupApp = new(RouterGroup)
