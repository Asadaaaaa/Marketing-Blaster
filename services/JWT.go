package services

import (
	"os"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(userId int) (string, error) {
	claims := jwt.MapClaims{}
	claims["userId"] = userId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_TOKEN_SECRET")))
}

func GenerateTokenAdmin(userId int, isAdmin bool) (string, error) {
	claims := jwt.MapClaims{}
	claims["userId"] = userId
	claims["isAdmin"] = isAdmin

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_TOKEN_SECRET")))
}
