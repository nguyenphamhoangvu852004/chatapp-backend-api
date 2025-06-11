package initialize

import (
	"chapapp-backend-api/global"
	"chapapp-backend-api/internal/router"

	"github.com/gin-contrib/cors"
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

	// CORS config
	config := cors.Config{
		AllowOrigins: []string{global.Config.Cors.Url},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
		// AllowCredentials: true,
		// MaxAge:           12 * time.Hour,
	}
	// r.Use() //corss
	r.Use(cors.New(config))

	// r.Use() //limiter globale

	authRouter := router.RouterGroupApp.AuthRouter
	profileRouter := router.RouterGroupApp.ProfileRouter
	accountRouter := router.RouterGroupApp.AccountRouter
	friendShipRouter := router.RouterGroupApp.FriendShipRouter
	mainGroup := r.Group("/api/v1")
	{
		mainGroup.GET("/checkStatus", func(c *gin.Context) { c.JSON(200, gin.H{"message": "ok"}) }) // tracking monitor
	}
	{
		authRouter.InitAuthRouter(mainGroup)
		profileRouter.InitProfileRouter(mainGroup)
		accountRouter.InitAccountRouter(mainGroup)
		friendShipRouter.InitFriendShipRouter(mainGroup)
	}

	return r
}
