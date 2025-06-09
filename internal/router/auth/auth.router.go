package auth

import (
	"chapapp-backend-api/internal/wire"

	"github.com/gin-gonic/gin"
)

type AuthRouter struct {
}

func (userRouter *AuthRouter) InitAuthRouter(router *gin.RouterGroup) {
	//public router
	authController, _ := wire.InitModuleAuth()
	userPublicRouter := router.Group("/auth")
	{
		userPublicRouter.POST("/sendOtp", authController.SendOTP)
		userPublicRouter.POST("/verifyOtp", authController.VerifyOTP)
		userPublicRouter.POST("/register", authController.Register)
		userPublicRouter.POST("/login", authController.Login)
	}

	//private router
	userPrivateRouter := router.Group("/user")
	// userPrivateRouter.Use(middleware.Limiter())
	// userPrivateRouter.Use(middleware.AuthMiddleware())
	// userPrivateRouter.Use(middleware.PermissionMiddleware())
	{
		userPrivateRouter.GET("/getInfo/:id")
	}

}
