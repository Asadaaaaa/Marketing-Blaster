package services

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
)

func GetPromotionTextFromPromptAIService(prompt string) interface{} {
	url := "https://api.openai.com/v1/chat/completions"
	apiKey := os.Getenv("OPENAI_API_KEY")
	requestBody := map[string]interface{}{
		"model": "gpt-3.5-turbo",
		"messages": []map[string]string{
			{
				"role":    "system",
				"content": "Saya adalah seorang marketing email blaster terbaik untuk memberikan kalimat promosi iklan yang menarik. Balas dengan menggunakan emoticon sehingga lebih menarik. dan sertakan link berdasarkan prompt",
			},
			{
				"role":    "user",
				"content": prompt,
			},
		},
		"temperature":       1,
		"max_tokens":        256,
		"top_p":             1,
		"frequency_penalty": 0,
		"presence_penalty":  0,
	}

	requestBodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		return -1
	}

	// Create the HTTP request.
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return -2
	}

	// Set the request headers.
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	// Send the HTTP request.
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return -3
	}
	defer resp.Body.Close()

	// Read the response body.
	var responseBody map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	if err != nil {
		return -4
	}

	content := responseBody["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)

	return content
}
