package repository

import (
	"fmt"
	"strings"

	"github.com/artem-shestakov/to-do/internal/models"
	"github.com/jmoiron/sqlx"
)

type TaskRepository struct {
	db *sqlx.DB
}

func NewTaskRepository(db *sqlx.DB) *TaskRepository {
	return &TaskRepository{
		db: db,
	}
}

func (r *TaskRepository) Create(listId int, task models.ToDoTask) (int, error) {
	// Create task query
	var taskId int
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description, list_id) VALUES ($1, $2, $3) RETURNING id", tasksTable)
	row := r.db.QueryRow(createListQuery, task.Title, task.Description, listId)
	err := row.Scan(&taskId)
	if err != nil {
		return 0, err
	}

	return taskId, nil
}

func (r *TaskRepository) GetTasks(listId int) ([]models.ToDoTask, error) {
	var tasks []models.ToDoTask

	query := fmt.Sprintf("SELECT * FROM %s WHERE list_id = $1", tasksTable)
	err := r.db.Select(&tasks, query, listId)
	return tasks, err
}

func (r *TaskRepository) GetTaskById(taskId, listId int) (models.ToDoTask, error) {
	var task models.ToDoTask

	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1 AND list_id = $2", tasksTable)
	err := r.db.Get(&task, query, taskId, listId)
	return task, err
}

func (r *TaskRepository) UpdateTask(taskId, listId int, task models.UpdateToDoTask) error {
	// Parse new values, create args for query and count args
	sets := make([]string, 0)
	args := make([]interface{}, 0)
	argID := 1

	if task.Title != "" {
		sets = append(sets, fmt.Sprintf("title=$%d", argID))
		args = append(args, task.Title)
		argID++
	}

	if task.Description != "" {
		sets = append(sets, fmt.Sprintf("description=$%d", argID))
		args = append(args, task.Description)
		argID++
	}

	// Update is_done param
	sets = append(sets, fmt.Sprintf("is_done=$%d", argID))
	args = append(args, task.IsDone)
	argID++

	setsQuery := strings.Join(sets, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d AND list_id = $%d",
		tasksTable,
		setsQuery,
		argID,
		argID+1,
	)

	fmt.Println(query)

	args = append(args, taskId, listId)
	result, err := r.db.Exec(query, args...)
	if err != nil {
		return err
	}

	// Check if update list
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return fmt.Errorf("task not found")
	}

	return err
}

func (r *TaskRepository) DeleteTask(taskId, listId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1 AND list_id = $2", tasksTable)
	result, err := r.db.Exec(query, taskId, listId)
	if err != nil {
		return err
	}

	// Check if delete list
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return fmt.Errorf("task not found")
	}

	return err
}
