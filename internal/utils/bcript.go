package utils

import (
	"chapapp-backend-api/global"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func getAccessSecretKey() []byte {
	return []byte(global.Config.Jwt.AccessTokenSecret)
}

func getRefreshSecretKey() []byte {
	return []byte(global.Config.Jwt.RefreshTokenSecret)
}

func getAccessTokenTTL() time.Duration {
	return time.Duration(global.Config.Jwt.AccessTokenExpiriedTime) * time.Second
}

func getRefreshTokenTTL() time.Duration {
	return time.Duration(global.Config.Jwt.RefreshTokenExpiriedTime) * time.Second
}

func GenerateToken(payload map[string]interface{}, secret []byte, ttl time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"exp": time.Now().Add(ttl).Unix(),
		"iat": time.Now().Unix(),
	}

	for k, v := range payload {
		claims[k] = v
	}

	// Tạo token và ký
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	fmt.Println(ttl)
	fmt.Println(secret)
	return token.SignedString(secret)
}

func GenerateAccessToken(userID uint, email string) (string, error) {
	payload := map[string]interface{}{
		"id":    userID,
		"email": email,
	}
	return GenerateToken(payload, getAccessSecretKey(), getAccessTokenTTL())
}

func GenerateRefreshToken(userID uint) (string, error) {
	payload := map[string]interface{}{
		"id":      userID,
	}
	return GenerateToken(payload, getRefreshSecretKey(), getRefreshTokenTTL())
}
