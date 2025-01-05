package token

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrMissingToken = errors.New("missing token")
	ErrInvalidToken = errors.New("invalid token")
)

// GenerateToken creates a JWT token for the specified userID and valid duration.
func GenerateToken(userID int64, ttl time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"iss": "donorbox",
			"id":  userID,
			"exp": time.Now().Add(ttl).Unix(),
		})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// VerifyToken parses a JWT token and returns the associated user ID.
func VerifyToken(tokenString string) (int64, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			slog.Debug("wrong method")
			return nil, ErrInvalidToken
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		slog.Debug(fmt.Sprintf("can not parse token: %v", err))
		return 0, ErrInvalidToken
	}

	if !token.Valid {
		slog.Debug("invalid token")
		return 0, ErrInvalidToken
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		panic("unable to decode claims")
	}

	id, ok := claims["id"]
	if !ok {
		panic("missing id claims")
	}

	return int64(id.(float64)), nil
}

// VerifyRequestToken grabs the JWT token from a request header and call VerifyToken to verifies it.
func VerifyRequestToken(r *http.Request) (int64, error) {
	val := r.Header.Get("Authorization")
	if val == "" {
		return 0, ErrMissingToken
	}

	parts := strings.Split(val, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return 0, ErrInvalidToken
	}

	tokenString := parts[1]

	id, err := VerifyToken(tokenString)
	if err != nil {
		return 0, ErrInvalidToken
	}

	return id, nil
}
