package controllers

import (
	"Marketing-Blaster/models"
	"Marketing-Blaster/request"
	"Marketing-Blaster/services"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func SendMailerController(ctx *fiber.Ctx) error {
	mailer := new(request.SendMailerRequest)
	if err := ctx.BodyParser(mailer); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Bad Request",
			"type":    "controller",
			"err": fiber.Map{
				"code": -1,
			},
		})
	}

	for _, item := range mailer.Email {
		services.SendMail(item, mailer.Subject, mailer.Body)
	}

	emails := strings.Join(mailer.Email, ", ")
	newLog := models.MailLog{
		Email:   emails,
		Subject: mailer.Subject,
		Body:    mailer.Body,
		// get user id from middleware
		UserID: int(ctx.Locals("userId").(float64)),
	}

	errCreateLog := models.DB.Create(&newLog).Error
	if errCreateLog != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Somthing went wrong with logger",
			"type":    "controller",
			"err": fiber.Map{
				"code": -2,
			},
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"message": "Success",
	})
}

func AIGetBodyController(ctx *fiber.Ctx) error {
	body := new(request.AIGetBodyMailerRequest)
	if err := ctx.BodyParser(body); err != nil {
		return err
	}

	promotionText := services.GetPromotionTextFromPromptAIService(body.Prompt)

	if promotionText == -1 || promotionText == -2 || promotionText == -3 || promotionText == -4 {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Somthing went wrong",
			"type":    "controller",
			"err": fiber.Map{
				"code": -1,
			},
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"message": "success",
		"data":    promotionText,
	})
}

func SendWithAIController(ctx *fiber.Ctx) error {
	// Request Body Validator
	body := new(request.SendWithAIMailerRequest)
	if err := ctx.BodyParser(body); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Bad Request",
			"type":    "controller",
			"err": fiber.Map{
				"code": -1,
			},
		})
	}

	promotionText := services.GetPromotionTextFromPromptAIService(body.Prompt)
	if promotionText == -1 || promotionText == -2 || promotionText == -3 || promotionText == -4 {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Somthing went wrong",
			"type":    "controller",
			"err": fiber.Map{
				"code": -2,
			},
		})
	}

	for _, item := range body.Email {
		services.SendMail(item, body.Subject, promotionText.(string))
	}

	emails := strings.Join(body.Email, ", ")
	newLog := models.MailLog{
		Email:   emails,
		Subject: body.Subject,
		Body:    promotionText.(string),
		// get user id from middleware
		UserID: int(ctx.Locals("userId").(float64)),
	}

	errCreateLog := models.DB.Create(&newLog).Error
	if errCreateLog != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Somthing went wrong with logger",
			"type":    "controller",
			"err": fiber.Map{
				"code": -3,
			},
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"message": "Success",
	})
}
