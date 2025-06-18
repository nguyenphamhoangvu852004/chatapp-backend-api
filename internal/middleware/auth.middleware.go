package middleware

import (
	"chapapp-backend-api/global"
	"chapapp-backend-api/pkg/response"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// func AuthMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		token := c.GetHeader("Authorization")
// 		if token != "valid-token" {
// 			response.ErrorReponse(c, http.StatusUnauthorized, "Unauthorized")
// 			c.Abort()
// 			return
// 		}
// 		payload, err := utils.ParseToken(token)
// 		if err != nil {
// 			response.ErrorReponse(c, http.StatusUnauthorized, "Unauthorized")
// 			c.Abort()
// 			return
// 		}
// 		fmt.Println(payload)

// 		c.Next()
// 	}
// }

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Lấy token từ header Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			response.ErrorReponse(c, http.StatusUnauthorized, "Missing or invalid Authorization header")
			c.Abort()
			return
		}

		// Cắt chuỗi "Bearer " ra để lấy token thật
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Parse và xác thực token
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			// Đảm bảo phương thức ký là HS256
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrTokenSignatureInvalid
			}
			return []byte(global.Config.Jwt.AccessTokenSecret), nil
		}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

		// Kiểm tra token hợp lệ và chưa hết hạn
		if err != nil || !token.Valid {
			response.ErrorReponse(c, http.StatusUnauthorized, "Invalid or expired token")
			c.Abort()
			return
		}

		// Lấy claims và gán vào context
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			exp, ok := claims["exp"].(float64)
			if ok && time.Now().Unix() > int64(exp) {
				response.ErrorReponse(c, http.StatusUnauthorized, "Token has expired")
				c.Abort()
				return
			}

			// Ví dụ gán userID vào context
			userID, _ := claims["id"]
			email, _ := claims["email"]
			roles, _ := claims["roles"]
			c.Set("userId", userID)
			c.Set("email", email)
			c.Set("roles", roles)
		} else {
			response.ErrorReponse(c, http.StatusUnauthorized, "Invalid token claims")
			c.Abort()
			return
		}

		c.Next()
	}
}
