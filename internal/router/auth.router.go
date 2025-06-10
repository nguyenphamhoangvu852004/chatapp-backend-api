package router

import (
	"chapapp-backend-api/internal/wire"

	"github.com/gin-gonic/gin"
)

type AuthRouter struct {
}

func (userRouter *AuthRouter) InitAuthRouter(router *gin.RouterGroup) {
	//public router
	authController, _ := wire.InitModuleAuth()
	authPublicRouter := router.Group("/auth")
	{
		authPublicRouter.POST("/sendOtp", authController.SendOTP)
		authPublicRouter.POST("/verifyOtp", authController.VerifyOTP)
		authPublicRouter.POST("/register", authController.Register)
		authPublicRouter.POST("/login", authController.Login)
		authPublicRouter.PUT("/resetPassword", authController.ResetPassword)
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
