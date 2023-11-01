package request

type AdminLoginAuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
