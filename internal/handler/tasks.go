package handler

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/artem-shestakov/to-do/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetUserList(c *gin.Context) (int, int, error) {
	// Check user and list list
	user_id, err := h.GetUserID(c)
	if err != nil {
		h.responseError(c, http.StatusUnauthorized, err, "user token is not valid")
		return 0, 0, err
	}

	// Get list id from url params
	id := c.Param("id")
	list_id, err := strconv.Atoi(id)
	if err != nil {
		h.responseError(c, http.StatusBadRequest, err, fmt.Sprintf("list id %s is invalid", id))
		return 0, 0, err
	}

	_, err = h.service.GetListById(user_id, list_id)
	if err != nil {
		if err == sql.ErrNoRows {
			h.responseError(c, http.StatusNotFound, err, fmt.Sprintf("list with id %d of user id %d is not found", list_id, user_id))
			return 0, 0, err
		}
		h.responseError(c, http.StatusInternalServerError, err, fmt.Sprintf("can't get list id %d of user id %d", list_id, user_id))
		return 0, 0, err
	}
	return user_id, list_id, nil
}

func (h *Handler) GetTasks(c *gin.Context) {
	// Get user's and list's ids
	user_id, list_id, err := h.GetUserList(c)
	if err != nil {
		return
	}

	// Get tasks
	tasks, err := h.service.ToDoTask.GetTasks(list_id)
	if err != nil {
		h.responseError(c, http.StatusInternalServerError, err, fmt.Sprintf("can't get tasks of user id %d user's list id %d", user_id, list_id))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"tasks": tasks,
	})
}

func (h *Handler) GetTask(c *gin.Context) {
	// Get user's and list's ids
	user_id, list_id, err := h.GetUserList(c)
	if err != nil {
		return
	}

	id := c.Param("task_id")
	task_id, err := strconv.Atoi(id)
	if err != nil {
		h.responseError(c, http.StatusBadRequest, err, fmt.Sprintf("task id %s is invalid", id))
		return
	}

	task, err := h.service.ToDoTask.GetTaskById(task_id, list_id)
	if err != nil {
		if err == sql.ErrNoRows {
			h.responseError(c, http.StatusNotFound, err, fmt.Sprintf("task with id %d of user id %d is not found", task_id, user_id))
			return
		}
		h.responseError(c, http.StatusInternalServerError, err, fmt.Sprintf("can't get list id %d of user id %d", task_id, user_id))
		return
	}

	// Return list
	c.JSON(http.StatusOK, task)
}

func (h *Handler) CreateTask(c *gin.Context) {
	var task models.ToDoTask

	// Get user's and list's ids
	_, list_id, err := h.GetUserList(c)
	if err != nil {
		return
	}

	// Parsing JSON
	if err := c.ShouldBindJSON(&task); err != nil {
		if err.Error() == "EOF" {
			err = errors.New("request body is empty")
		}
		h.responseError(c, http.StatusBadRequest, err, "json parsing error")
		return
	}

	task_id, err := h.service.ToDoTask.Create(list_id, task)
	if err != nil {
		h.responseError(c, http.StatusInternalServerError, err, "can't create task")
		return
	}

	// Return created list id
	c.JSON(http.StatusCreated, gin.H{
		"id": task_id,
	})
}

func (h *Handler) UpdateTask(c *gin.Context) {
	// Get user's and list's ids
	user_id, list_id, err := h.GetUserList(c)
	if err != nil {
		return
	}

	id := c.Param("task_id")
	task_id, err := strconv.Atoi(id)
	if err != nil {
		h.responseError(c, http.StatusBadRequest, err, fmt.Sprintf("task id %s is invalid", id))
		return
	}

	var task models.UpdateToDoTask

	// Parsing JSON
	if err := c.ShouldBindJSON(&task); err != nil {
		if err.Error() == "EOF" {
			err = errors.New("request body is empty")
		}
		h.responseError(c, http.StatusBadRequest, err, "json parsing error")
		return
	}

	err = h.service.ToDoTask.UpdateTask(task_id, list_id, task)
	if err != nil {
		if strings.Contains(err.Error(), "task not found") {
			h.responseError(c, http.StatusNotFound, err, fmt.Sprintf("task with id %d of user id %d is not found", task_id, user_id))
			return
		}
		h.responseError(c, http.StatusInternalServerError, err, fmt.Sprintf("can't update task id %d of user id %d", task_id, user_id))
		return
	}

	// Return list
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func (h *Handler) DeleteTask(c *gin.Context) {
	// Get user's and list's ids
	user_id, list_id, err := h.GetUserList(c)
	if err != nil {
		return
	}

	// Get task id from URL
	id := c.Param("task_id")
	task_id, err := strconv.Atoi(id)
	if err != nil {
		h.responseError(c, http.StatusBadRequest, err, fmt.Sprintf("task id %s is invalid", id))
		return
	}

	err = h.service.ToDoTask.DeleteTask(task_id, list_id)
	if err != nil {
		if strings.Contains(err.Error(), "task not found") {
			h.responseError(c, http.StatusNotFound, err, fmt.Sprintf("task with id %d of user id %d is not found", task_id, user_id))
			return
		}
		h.responseError(c, http.StatusInternalServerError, err, fmt.Sprintf("can't delete task id %d of user id %d", task_id, user_id))
		return
	}

	// Return list
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
