package router

import (
	"chapapp-backend-api/internal/router/auth"
)

type RouterGroup struct {
	AuthRouter auth.AuthRouter
}

var RouterGroupApp = new(RouterGroup)
