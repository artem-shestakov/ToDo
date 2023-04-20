package models

type User struct {
	ID        int    `json:"id" db:"id"`
	FirstName string `json:"first_name" db:"first_name" binding:"required"`
	LasteName string `json:"last_name" db:"last_name" binding:"required"`
	Email     string `json:"email" db:"email" binding:"required,email"`
	Password  string `json:"password" db:"password" binding:"required"`
}
