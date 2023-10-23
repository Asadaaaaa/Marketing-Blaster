package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func AuthorizationMiddleware(c *fiber.Ctx) error {
	// Get token from header Authorizatioin
	tokenString := c.Get("Authorization")

	// Parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.ErrUnauthorized
		}

		// Return secret key
		return []byte("salman-marketing-blaster"), nil
	})

	// Check error
	if err != nil {
		return fiber.ErrUnauthorized
	}

	// Get claims
	claims := token.Claims.(jwt.MapClaims)

	// Set userId in locals
	c.Locals("userId", claims["userId"])

	// Continue stack
	return c.Next()
}
