package service

import (
	"github.com/artem-shestakov/to-do/internal/models"
	"github.com/artem-shestakov/to-do/internal/repository"
)

type Auth interface {
	CreateUser(user models.User) (int, error)
	GetUser(email, password string) (models.User, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(tokenString string) (int, error)
}

type ToDoList interface {
	Create(userId int, list models.ToDoList) (int, error)
	GetLists(userId int) ([]models.ToDoList, error)
	GetListById(userId, listId int) (models.ToDoList, error)
	UpdateList(userId, listId int, list models.UpdateToDoList) error
	DeleteList(userId, listId int) error
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
		Auth:     NewAuthService(repo),
		ToDoList: NewListService(repo),
	}
}
