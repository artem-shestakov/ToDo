package repository

import (
	"github.com/artem-shestakov/to-do/internal/models"
	"github.com/jmoiron/sqlx"
)

const (
	userTable = "users"
)

type Auth interface {
	CreateUser(user models.User) (int, error)
}

type ToDoList interface {
}

type ToDoTask interface {
}

type Repository struct {
	Auth
	ToDoList
	ToDoTask
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Auth: NewAuthRepository(db),
	}
}
