package models

import (
	"errors"
	"reflect"
)

type ToDoList struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}

type UpdateToDoList struct {
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
}

type ToDoTask struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsDone      bool   `json:"is_done"`
}

func (i UpdateToDoList) Validate() error {
	if reflect.ValueOf(i).IsZero() {
		return errors.New("receive data are empty or not correct")
	}
	return nil
}
