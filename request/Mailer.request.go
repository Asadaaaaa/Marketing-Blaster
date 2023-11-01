package request

type SendMailerRequest struct {
	Email   []string `json:"email"`
	Subject string   `json:"subject"`
	Body    string   `json:"body"`
}

type AIGetBodyMailerRequest struct {
	Prompt string `json:"prompt"`
}

type SendWithAIMailerRequest struct {
	Prompt  string   `json:"prompt"`
	Subject string   `json:"subject"`
	Email   []string `json:"email"`
}
