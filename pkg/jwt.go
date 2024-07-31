package pkg

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt"
)

type TokenClaims struct {
	ID        uint  `json:"id"`
	ExpiresAt int64 `json:"exp"`
	jwt.StandardClaims
}

func GenerateToken(key string, claims TokenClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secretKey := []byte(key)

	signedString, err := token.SignedString(secretKey)

	return signedString, err
}

func ValidateToken(authToken string, key string) (uint, error) {
	token, err := jwt.Parse(authToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(key), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok && !token.Valid {
		return 0, errors.New("unauthorized access")
	}

	userIdFloat := claims["id"].(float64)
	userId := uint(userIdFloat)
	return userId, nil

}
