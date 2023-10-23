package controllers

import (
	"Marketing-Blaster/request"
	"Marketing-Blaster/services"

	"github.com/gofiber/fiber/v2"

	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func SendMailerController(ctx *fiber.Ctx) error {
	mailer := new(request.SendMailerRequest)
	if err := ctx.BodyParser(mailer); err != nil {
		return err
	}

	for _, item := range mailer.Email {
		services.SendMail(item, mailer.Subject, mailer.Body)
	}

	return ctx.Status(200).JSON(fiber.Map{
		"message": "success",
	})
}

func AIGetBodyController(ctx *fiber.Ctx) error {
	body := new(request.AIGetBodyMailerRequest)
	if err := ctx.BodyParser(body); err != nil {
		return err
	}

	url := "https://api.openai.com/v1/chat/completions"
	apiKey := "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"
	requestBody := map[string]interface{}{
		"model": "gpt-3.5-turbo",
		"messages": []map[string]string{
			{
				"role":    "system",
				"content": "Saya adalah seorang marketing email blaster terbaik untuk memberikan kalimat promosi iklan yang menarik. Balas dengan menggunakan emoticon sehingga lebih menarik. dan sertakan link berdasarkan prompt",
			},
			{
				"role":    "user",
				"content": body.Prompt,
			},
		},
		"temperature":       1,
		"max_tokens":        256,
		"top_p":             1,
		"frequency_penalty": 0,
		"presence_penalty":  0,
	}

	// Convert the request body to JSON.
	requestBodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		panic(err)
	}

	// Create the HTTP request.
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		panic(err)
	}

	// Set the request headers.
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	// Send the HTTP request.
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Read the response body.
	var responseBody map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	if err != nil {
		panic(err)
	}

	// Print the response body.
	fmt.Println(responseBody)

	content := responseBody["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)

	return ctx.Status(200).JSON(fiber.Map{
		"message": "success",
		"data":    content,
	})
}
