package util

import (
	"context"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(ctx context.Context, userID string, exp int64, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     exp,
	})
	return token.SignedString([]byte(secret))
}

func ParseJWT(ctx context.Context, tokenStr string, secret string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if id, ok := claims["id"].(string); ok {
			return id, nil
		}

		return "", fmt.Errorf("id claim not found or invalid")
	}

	return "", fmt.Errorf("invalid token")
}
