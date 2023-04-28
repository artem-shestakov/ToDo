package service

import (
	"github.com/artem-shestakov/to-do/internal/models"
	"github.com/artem-shestakov/to-do/internal/repository"
)

type TaskService struct {
	repo *repository.Repository
}

func NewTaskService(repo *repository.Repository) *TaskService {
	return &TaskService{
		repo: repo,
	}
}

func (s *TaskService) Create(listId int, task models.ToDoTask) (int, error) {
	return s.repo.ToDoTask.Create(listId, task)
}

func (s *TaskService) GetTasks(listId int) ([]models.ToDoTask, error) {
	return s.repo.ToDoTask.GetTasks(listId)
}

func (s *TaskService) GetTaskById(taskId, listId int) (models.ToDoTask, error) {
	return s.repo.ToDoTask.GetTaskById(taskId, listId)
}

func (s *TaskService) UpdateTask(taskId, listId int, task models.UpdateToDoTask) error {
	if err := task.Validate(); err != nil {
		return err
	}
	return s.repo.ToDoTask.UpdateTask(taskId, listId, task)
}

func (s *TaskService) DeleteTask(taskId, listId int) error {
	return s.repo.ToDoTask.DeleteTask(taskId, listId)
}
