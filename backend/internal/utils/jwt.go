package utils

import (
	"backend/internal/dtos"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))

type Claims struct {
	User dtos.LoginUserDTO
	jwt.RegisteredClaims
}

func GenerateJWT(user dtos.LoginUserDTO) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		User: user,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtKey)
}

func ValidateJWT(tokenStr string) (dtos.LoginUserDTO, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return dtos.LoginUserDTO{}, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return dtos.LoginUserDTO{}, fmt.Errorf("invalid token")
	}

	return claims.User, nil
}
