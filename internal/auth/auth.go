package auth

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrInvalidSigningMethod = errors.New("invalid signing method")
	ErrInvalidToken         = errors.New("invalid token")
)

func GenerateToken(userID string, userType string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"iss": "donorbox",
			"id":  userID,
			"exp": time.Now().Add(72 * time.Hour),
		})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			slog.Debug("wrong method")
			return nil, ErrInvalidSigningMethod
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		slog.Debug(fmt.Sprintf("can not parse token: %v", err))
		return "", ErrInvalidToken
	}

	if !token.Valid {
		slog.Debug("invalid token")
		return "", ErrInvalidToken
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		panic("unable to decode claims")
	}

	id, ok := claims["id"]
	if !ok {
		panic("missing id claims")
	}

	return id.(string), nil
}
