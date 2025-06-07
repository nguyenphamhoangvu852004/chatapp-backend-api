package initialize

import (
	"chapapp-backend-api/global"
	"chapapp-backend-api/internal/router"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	var r *gin.Engine
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}

	// middleware

	// r.Use() //loggin
	// r.Use() //corss
	// r.Use() //limiter globale

	mananagerRouter := router.RouterGroupApp.ManagerRouter
	userRouter := router.RouterGroupApp.UserRouter

	mainGroup := r.Group("/api/v1")
	{
		mainGroup.GET("/checkStatus") // tracking monitor
	}
	{
		userRouter.UserRouter.InitUserRouter(mainGroup)
		userRouter.ProductRouter.InitProductRouter(mainGroup)
	}
	{
		mananagerRouter.UserRouter.InitUserRouter(mainGroup)
		mananagerRouter.AdminRouter.InitAdminRouter(mainGroup)
	}

	return r
}
