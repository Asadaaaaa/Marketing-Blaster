package models

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey;autoIncrement"` // Auto increment
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
