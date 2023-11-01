package request

// Register Validator body have email, password, and name
type RegisterAuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

// Login Validator body have email, password
type LoginAuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
