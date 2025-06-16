package router

import (
	"chapapp-backend-api/internal/middleware"
	"chapapp-backend-api/internal/wire"
	"net/http"

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
		authPublicRouter.GET("/validateToken", middleware.AuthMiddleware(), func(c *gin.Context) {
			userId, _ := c.Get("userId")
			mail, _ := c.Get("email")
			roles, _ := c.Get("roles")
			if userId == nil || mail == nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID or Email not found in context"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"userId": userId, "email": mail, "roles": roles, "isValid": true})
		})
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
