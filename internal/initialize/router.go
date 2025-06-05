package initialize

import (
	"chapapp-backend-api/internal/controller"
	"chapapp-backend-api/internal/middleware"
	"fmt"

	"github.com/gin-gonic/gin"
)

func AA() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("AAA ne")
		c.Next()
		fmt.Println("Quay lai lam cai con lai trong FUNC AA")
	}
}

func BB() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("BBB ne")
		c.Next()
		fmt.Println("Quay lai lam cai con lai trong FUNC BB")
	}
}

func CC(c *gin.Context) {
	fmt.Println("CCC ne")
	c.Next()
	fmt.Println("Quay lai lam cai con lai trong FUNC CC")
}

func InitRouter() *gin.Engine {

	r := gin.Default()
	r.Use(middleware.AuthMiddleware())
	v1 := r.Group("/v1")
	{
		v1.GET("/ping", controller.NewPongController().Pong)
		v1.GET("/user/:id", controller.NewUserController().GetUserById)
		// v1.POST("/ping", Pong)
		// v1.DELETE("/ping", Pong)
		// v1.PATCH("/ping", Pong)
		// v1.PUT("/ping", Pong)
		// v1.HEAD("/ping", Pong)
	}

	// v2 := r.Group("/v2")
	{
		// v2.GET("/ping", Pong)
		// v2.POST("/ping", Pong)
		// v2.DELETE("/ping", Pong)
		// v2.PATCH("/ping", Pong)
		// v2.PUT("/ping", Pong)
		// v2.HEAD("/ping", Pong)
	}

	return r
}
