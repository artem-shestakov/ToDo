package models

type User struct {
	ID        int    `json:"user_id" db:"user_id" binding:"required"`
	FirstName string `json:"first_name" db:"first_name" binding:"required"`
	LasteName string `json:"last_name" db:"last_name" binding:"required"`
}
