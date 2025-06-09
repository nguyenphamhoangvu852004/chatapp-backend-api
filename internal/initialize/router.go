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

	userRouter := router.RouterGroupApp.AuthRouter

	mainGroup := r.Group("/api/v1")
	{
		mainGroup.GET("/checkStatus", func(c *gin.Context) { c.JSON(200, gin.H{"message": "ok"}) }) // tracking monitor
	}
	{
		userRouter.InitAuthRouter(mainGroup)
	}

	return r
}
