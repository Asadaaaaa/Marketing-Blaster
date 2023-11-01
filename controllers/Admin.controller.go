package controllers

import (
	"Marketing-Blaster/models"
	"Marketing-Blaster/request"
	"Marketing-Blaster/services"

	"github.com/gofiber/fiber/v2"
)

func AdminLoginAuthController(ctx *fiber.Ctx) error {
	body := new(request.AdminLoginAuthRequest)

	if err := ctx.BodyParser(body); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Bad Request",
			"type":    "controller",
			"err": fiber.Map{
				"code": -1,
			},
		})
	}

	var adminModelData models.Admin
	err := models.DB.Where("username = ?", body.Username).First(&adminModelData).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "Username or Password is wrong",
			"type":    "controller",
			"err": fiber.Map{
				"code": -1,
			},
		})
	}

	if adminModelData.Password != body.Password {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "Username or Password is wrong",
			"type":    "controller",
			"err": fiber.Map{
				"code": -1,
			},
		})
	}

	token, err := services.GenerateTokenAdmin(int(adminModelData.ID), true)
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

func AdminGetAllUserController(ctx *fiber.Ctx) error {
	var users []models.User
	err := models.DB.Find(&users).Error
	if err != nil {
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
		"data":    users,
	})
}

func AdminGetAllMailLogsController(ctx *fiber.Ctx) error {
	var mailLogs []models.MailLog
	err := models.DB.Find(&mailLogs).Error
	if err != nil {
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
		"data":    mailLogs,
	})
}
