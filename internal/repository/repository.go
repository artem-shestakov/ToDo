package repository

import (
	"github.com/artem-shestakov/to-do/internal/models"
	"github.com/jmoiron/sqlx"
)

const (
	userTable      = "users"
	listsTable     = "lists"
	usersListTable = "users_lists"
)

type Auth interface {
	CreateUser(user models.User) (int, error)
	GetUser(email, password string) (models.User, error)
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

type Repository struct {
	Auth
	ToDoList
	ToDoTask
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Auth:     NewAuthRepository(db),
		ToDoList: NewListRepository(db),
	}
}
