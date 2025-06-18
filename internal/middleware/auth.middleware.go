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

//			c.Next()
//		}
//	}
type Role struct {
	Rolename string `json:"rolename"`
}

type UserPayload struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Roles    []Role `json:"roles"`
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			response.ErrorReponse(c, http.StatusUnauthorized, "Missing or invalid Authorization header")
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrTokenSignatureInvalid
			}
			return []byte(global.Config.Jwt.AccessTokenSecret), nil
		}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

		if err != nil || !token.Valid {
			response.ErrorReponse(c, http.StatusUnauthorized, "Invalid or expired token")
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			exp, _ := claims["exp"].(float64)
			if time.Now().Unix() > int64(exp) {
				response.ErrorReponse(c, http.StatusUnauthorized, "Token has expired")
				c.Abort()
				return
			}

			roles := []Role{}
			if rawRoles, ok := claims["roles"].([]interface{}); ok {
				for _, r := range rawRoles {
					if roleMap, ok := r.(map[string]interface{}); ok {
						if rolename, ok := roleMap["rolename"].(string); ok {
							roles = append(roles, Role{Rolename: rolename})
						}
					}
				}
			}

			user := map[string]interface{}{
				"id":       claims["id"],
				"email":    claims["email"],
				"username": claims["username"],
				"roles":    claims["roles"], // giữ nguyên, không ép kiểu slice<struct>
			}
			c.Set("user", user)
		} else {
			response.ErrorReponse(c, http.StatusUnauthorized, "Invalid token claims")
			c.Abort()
			return
		}

		c.Next()
	}
}

func VerifyRole(requiredRoles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userData, exists := c.Get("user")
		if !exists {
			response.ErrorReponse(c, http.StatusUnauthorized, "No user in context")
			return
		}

		userMap, ok := userData.(map[string]interface{})
		if !ok {
			response.ErrorReponse(c, http.StatusInternalServerError, "User data format error")
			return
		}

		rolesRaw, ok := userMap["roles"].([]interface{})
		if !ok {
			response.ErrorReponse(c, http.StatusForbidden, "Invalid roles format")
			return
		}

		userRoles := make([]string, 0)
		for _, r := range rolesRaw {
			if roleStr, ok := r.(string); ok {
				userRoles = append(userRoles, roleStr)
			}
		}

		// Check if user has at least one required role
		hasPermission := false
		for _, userRole := range userRoles {
			for _, required := range requiredRoles {
				if userRole == required {
					hasPermission = true
					break
				}
			}
			if hasPermission {
				break
			}
		}

		if !hasPermission {
			response.ErrorReponse(c, http.StatusForbidden, "Forbidden: No permission")
			return
		}

		c.Next()
	}
}
