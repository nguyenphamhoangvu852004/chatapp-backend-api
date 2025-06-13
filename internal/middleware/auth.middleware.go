package middleware

import (
	"chapapp-backend-api/internal/utils"
	"chapapp-backend-api/pkg/response"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != "valid-token" {
			response.ErrorReponse(c, http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		payload, err := utils.ParseToken(token)
		if err != nil {
			response.ErrorReponse(c, http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		fmt.Println(payload)

		c.Next()
	}
}
