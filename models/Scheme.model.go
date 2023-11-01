package models

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey;autoIncrement"` // Auto increment
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Mail Logs
type MailLog struct {
	ID      uint   `json:"id" gorm:"primaryKey;autoIncrement"` // Auto increment
	UserID  int    `json:"user_id"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
	Email   string `json:"email"`
}

type Admin struct {
	ID       uint   `json:"id" gorm:"primaryKey;autoIncrement"` // Auto increment
	Username string `json:"username"`
	Password string `json:"password"`
}
