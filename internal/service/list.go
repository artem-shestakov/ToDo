package service

import (
	"github.com/artem-shestakov/to-do/internal/models"
	"github.com/artem-shestakov/to-do/internal/repository"
)

type ListService struct {
	repo *repository.Repository
}

func NewListService(repo *repository.Repository) *ListService {
	return &ListService{
		repo: repo,
	}
}

func (s *ListService) Create(userId int, list models.ToDoList) (int, error) {
	return s.repo.ToDoList.Create(userId, list)
}

func (s *ListService) GetLists(userId int) ([]models.ToDoList, error) {
	return s.repo.ToDoList.GetLists(userId)
}

func (s *ListService) GetListById(userId, listId int) (models.ToDoList, error) {
	return s.repo.ToDoList.GetListById(userId, listId)
}

func (s *ListService) UpdateList(userId, listId int, list models.UpdateToDoList) error {
	if err := list.Validate(); err != nil {
		return err
	}
	return s.repo.ToDoList.UpdateList(userId, listId, list)
}

func (s *ListService) DeleteList(userId, listId int) error {
	return s.repo.ToDoList.DeleteList(userId, listId)
}
