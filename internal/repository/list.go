package repository

import (
	"fmt"
	"strings"

	"github.com/artem-shestakov/to-do/internal/models"
	"github.com/jmoiron/sqlx"
)

type ListRepository struct {
	db *sqlx.DB
}

func NewListRepository(db *sqlx.DB) *ListRepository {
	return &ListRepository{
		db: db,
	}
}

func (r *ListRepository) Create(userId int, list models.ToDoList) (int, error) {
	// Create transaction
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	// Create list query
	var listId int
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", listsTable)
	row := tx.QueryRow(createListQuery, list.Title, list.Description)
	err = row.Scan(&listId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	// Add row into users_lists table
	createUserListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", usersListTable)
	_, err = tx.Exec(createUserListQuery, userId, listId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return listId, tx.Commit()
}

func (r *ListRepository) GetLists(userId int) ([]models.ToDoList, error) {
	var lists []models.ToDoList

	query := fmt.Sprintf("SELECT l.* FROM %s l INNER JOIN %s ul on l.id = ul.list_id WHERE ul.user_id = $1", listsTable, usersListTable)
	err := r.db.Select(&lists, query, userId)
	return lists, err
}

func (r *ListRepository) GetListById(userId, listId int) (models.ToDoList, error) {
	var list models.ToDoList

	query := fmt.Sprintf("SELECT l.* FROM %s l INNER JOIN %s ul on l.id = ul.list_id WHERE ul.user_id = $1 AND ul.list_id = $2", listsTable, usersListTable)
	err := r.db.Get(&list, query, userId, listId)
	return list, err
}

func (r *ListRepository) UpdateList(userId, listId int, list models.UpdateToDoList) error {
	// Parse new values, create args for query and count args
	sets := make([]string, 0)
	args := make([]interface{}, 0)
	argID := 1

	if list.Title != "" {
		sets = append(sets, fmt.Sprintf("title=$%d", argID))
		args = append(args, list.Title)
		argID++
	}

	if list.Description != "" {
		sets = append(sets, fmt.Sprintf("description=$%d", argID))
		args = append(args, list.Description)
		argID++
	}
	setsQuery := strings.Join(sets, ", ")

	query := fmt.Sprintf("UPDATE %s l SET %s FROM %s ul WHERE l.id = ul.list_id AND ul.user_id = $%d AND ul.list_id = $%d",
		listsTable,
		setsQuery,
		usersListTable,
		argID,
		argID+1,
	)

	fmt.Printf("UPDATE %s l SET %s FROM %s ul WHERE l.id = ul.list_id AND ul.user_id = $%d AND ul.list_id = $%d",
		listsTable,
		setsQuery,
		usersListTable,
		argID,
		argID+1,
	)

	args = append(args, userId, listId)
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
		return fmt.Errorf("list not found")
	}

	return err
	// e := reflect.ValueOf(&list).Elem()

	// for i := 0; i < e.NumField(); i++ {
	// 	fieldName := e.Type().Field(i).Tag.Get("db")
	// 	fieldType := e.Type().Field(i).Type
	// 	fieldValue := e.Field(i).Interface()

	// 	fmt.Printf("%v %v %v\n", fieldName, fieldType, fieldValue)
	// }
}

func (r *ListRepository) DeleteList(userId, listId int) error {
	query := fmt.Sprintf("DELETE FROM %s l USING %s ul WHERE l.id = ul.list_id AND ul.user_id = $1 AND ul.list_id = $2", listsTable, usersListTable)
	result, err := r.db.Exec(query, userId, listId)
	if err != nil {
		return err
	}

	// Check if delete list
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return fmt.Errorf("list not found")
	}

	return err
}
