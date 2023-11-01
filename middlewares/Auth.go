package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"

	"os"
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
		return []byte(os.Getenv("JWT_TOKEN_SECRET")), nil
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

func AuthorizationMiddlewareAdmin(c *fiber.Ctx) error {
	// Get token from header Authorizatioin
	tokenString := c.Get("Authorization")

	// Parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.ErrUnauthorized
		}

		// Return secret key
		return []byte(os.Getenv("JWT_TOKEN_SECRET")), nil
	})

	// Check error
	if err != nil {
		return fiber.ErrUnauthorized
	}

	// Get claims
	claims := token.Claims.(jwt.MapClaims)

	// Set userId in locals
	c.Locals("userId", claims["userId"])
	c.Locals("isAdmin", claims["isAdmin"])
	if claims["isAdmin"] == false {
		return c.Status(400).JSON(fiber.Map{
			"message": "Bad Request",
			"type":    "authorization",
			"err": fiber.Map{
				"code": -1,
			},
		})
	}

	// Continue stack
	return c.Next()
}
