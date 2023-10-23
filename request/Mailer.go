package request

type SendMailerRequest struct {
	Email   []string `json:"email"`
	Subject string   `json:"subject"`
	Body    string   `json:"body"`
}

type AIGetBodyMailerRequest struct {
	Prompt string `json:"prompt"`
}
