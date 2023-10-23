package services

import (
	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(userId int) (string, error) {
	claims := jwt.MapClaims{}
	claims["userId"] = userId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("salman-marketing-blaster"))
}
