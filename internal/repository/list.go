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
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description, user_id) VALUES ($1, $2, $3) RETURNING id", listsTable)
	row := tx.QueryRow(createListQuery, list.Title, list.Description, userId)
	err = row.Scan(&listId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	// // Add row into users_lists table
	// createUserListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", usersListTable)
	// _, err = tx.Exec(createUserListQuery, userId, listId)
	// if err != nil {
	// 	tx.Rollback()
	// 	return 0, err
	// }

	return listId, tx.Commit()
}

func (r *ListRepository) GetLists(userId int) ([]models.ToDoList, error) {
	var lists []models.ToDoList

	// query := fmt.Sprintf("SELECT l.* FROM %s l INNER JOIN %s ul on l.id = ul.list_id WHERE ul.user_id = $1", listsTable, usersListTable)
	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id = $1", listsTable)
	err := r.db.Select(&lists, query, userId)
	return lists, err
}

func (r *ListRepository) GetListById(userId, listId int) (models.ToDoList, error) {
	var list models.ToDoList

	// query := fmt.Sprintf("SELECT l.* FROM %s l INNER JOIN %s ul on l.id = ul.list_id WHERE ul.user_id = $1 AND ul.list_id = $2", listsTable, usersListTable)
	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id = $1 AND id = $2", listsTable)
	err := r.db.Get(&list, query, userId, listId)
	return list, err
}

func (r *ListRepository) GetListByTitle(userId int, listTitle string) (models.ToDoList, error) {
	var list models.ToDoList

	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id = $1 AND title = $2", listsTable)
	err := r.db.Get(&list, query, userId, listTitle)
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

	query := fmt.Sprintf("UPDATE %s SET %s WHERE user_id = $%d AND id = $%d",
		listsTable,
		setsQuery,
		argID,
		argID+1,
	)

	fmt.Println(query)

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
	query := fmt.Sprintf("DELETE FROM %s WHERE user_id = $1 AND id = $2", listsTable)
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
