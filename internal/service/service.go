package service

import (
	"github.com/artem-shestakov/to-do/internal/models"
	"github.com/artem-shestakov/to-do/internal/repository"
)

type Auth interface {
	CreateUser(user models.User) (int, error)
}

type ToDoList interface {
}

type ToDoTask interface {
}

type Service struct {
	Auth
	ToDoList
	ToDoTask
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Auth: NewAuthService(repo),
	}
}
