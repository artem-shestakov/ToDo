package models

import (
	"errors"
	"reflect"
)

type ToDoList struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
	UserId      int    `json:"user_id" db:"user_id"`
}

type UpdateToDoList struct {
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
}

type ToDoTask struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
	IsDone      bool   `json:"is_done" db:"is_done"`
	ListId      int    `json:"list_id" db:"list_id"`
}

type UpdateToDoTask struct {
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	IsDone      bool   `json:"is_done" db:"is_done"`
}

func (i UpdateToDoList) Validate() error {
	if reflect.ValueOf(i).IsZero() {
		return errors.New("receive data are empty or not correct")
	}
	return nil
}

func (i UpdateToDoTask) Validate() error {
	if reflect.ValueOf(i).IsZero() {
		return errors.New("receive data are empty or not correct")
	}
	return nil
}
