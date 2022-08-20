package util

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"go-coinstream/pkg/core/entity"
	"os"
	"time"
)

var SecretKey = []byte(os.Getenv("JWT_SECRET"))
var ExpirationTime = 15 * time.Minute

func GenerateToken(user entity.User) (string, error) {
	expirationTime := time.Now().Add(ExpirationTime)

	claims := jwt.RegisteredClaims{
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(expirationTime),
		Subject:   user.Username,
		Issuer:    "Coinstream",
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(SecretKey)
	if err != nil {
		return "", err
	}

	return token, nil
}

func ValidateToken(token string) (*jwt.RegisteredClaims, error) {
	var claims jwt.RegisteredClaims

	jwtToken, err := jwt.ParseWithClaims(
		token,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			return SecretKey, nil
		},
	)

	if !jwtToken.Valid {
		if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
			return nil, errors.New("token expired or not valid yet")
		}

	}
	return &claims, nil
}
