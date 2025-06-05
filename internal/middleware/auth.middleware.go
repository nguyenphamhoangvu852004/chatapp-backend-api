package middleware

import (
	"chapapp-backend-api/pkg/response"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != "valid-token" {
			response.ErrorReponse(c, 20003, "Loi roiii neee")
			c.Abort()
			return
		}
		c.Next()
	}
}
