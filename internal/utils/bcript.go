package utils

import (
	"chapapp-backend-api/global"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Cấu trúc payload mà Vũ lưu trong token (ví dụ: userId và email)
type Claims struct {
	UserID uint   `json:"id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

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

func GenerateAccessToken(userID uint, email string, roles []string) (string, error) {
	payload := map[string]interface{}{
		"id":    userID,
		"email": email,
		"roles": roles,
	}
	return GenerateToken(payload, getAccessSecretKey(), getAccessTokenTTL())
}

func GenerateRefreshToken(userID uint) (string, error) {
	payload := map[string]interface{}{
		"id": userID,
	}
	return GenerateToken(payload, getRefreshSecretKey(), getRefreshTokenTTL())
}

// Hàm parse token
func ParseToken(tokenStr string) (*Claims, error) {
	secret := []byte(global.Config.Jwt.AccessTokenSecret) // ← lấy từ env hoặc config

	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (any, error) {
		// Đảm bảo dùng đúng signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
