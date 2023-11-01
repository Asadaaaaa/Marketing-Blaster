package controllers

import (
	"Marketing-Blaster/models"
	"Marketing-Blaster/request"
	"Marketing-Blaster/services"

	"github.com/gofiber/fiber/v2"
)

func RegisterAuthController(ctx *fiber.Ctx) error {
	user := new(request.RegisterAuthRequest)
	if err := ctx.BodyParser(user); err != nil {
		return err
	}

	newUser := models.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}

	errCreateUser := models.DB.Create(&newUser).Error
	if errCreateUser != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Somthing went wrong",
			"type":    "controller",
			"err": fiber.Map{
				"code": -1,
			},
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"message": "Success",
	})
}

func LoginAuthController(ctx *fiber.Ctx) error {
	user := new(request.LoginAuthRequest)

	if err := ctx.BodyParser(user); err != nil {
		return err
	}

	var newUser models.User
	err := models.DB.Where("email = ?", user.Email).First(&newUser).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "Email or Password is wrong",
			"type":    "controller",
			"err": fiber.Map{
				"code": -1,
			},
		})
	}

	if newUser.Password != user.Password {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "Email or Password is wrong",
			"type":    "controller",
			"err": fiber.Map{
				"code": -1,
			},
		})
	}

	token, err := services.GenerateToken(int(newUser.ID))
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Somthing went wrong with generate token",
			"type":    "controller",
			"err": fiber.Map{
				"code": -2,
			},
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"message": "Success",
		"data": fiber.Map{
			"token": token,
		},
	})
}
