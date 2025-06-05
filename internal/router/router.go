package router

import (
	c "chapapp-backend-api/internal/controller"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {

	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/ping", c.NewPongController().Pong)
		v1.GET("/user/:id", c.NewUserController().GetUserById)
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
