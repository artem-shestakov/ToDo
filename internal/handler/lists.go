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

func (h *Handler) GetLists(c *gin.Context) {
	// Get user from JWT
	user_id, err := h.GetUserID(c)
	if err != nil {
		h.responseError(c, http.StatusUnauthorized, err, "user token is not valid")
		return
	}

	lists, err := h.service.ToDoList.GetLists(user_id)
	if err != nil {
		h.responseError(c, http.StatusInternalServerError, err, fmt.Sprintf("can't get lists of user id %d", user_id))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"lists": lists,
	})
}

func (h *Handler) GetList(c *gin.Context) {
	// Get user from JWT
	user_id, err := h.GetUserID(c)
	if err != nil {
		h.responseError(c, http.StatusUnauthorized, err, "user token is not valid")
		return
	}

	// Get list id from url params
	id := c.Param("id")
	list_id, err := strconv.Atoi(id)
	if err != nil {
		h.responseError(c, http.StatusBadRequest, err, fmt.Sprintf("list id %s is invalid", id))
		return
	}

	// Get list
	list, err := h.service.ToDoList.GetListById(user_id, list_id)
	if err != nil {
		if err == sql.ErrNoRows {
			h.responseError(c, http.StatusNotFound, err, fmt.Sprintf("list with id %d of user id %d is not found", list_id, user_id))
			return
		}
		h.responseError(c, http.StatusInternalServerError, err, fmt.Sprintf("can't get list id %d of user id %d", list_id, user_id))
		return
	}

	// Return list
	c.JSON(http.StatusOK, list)
}

func (h *Handler) CreateList(c *gin.Context) {
	var list models.ToDoList

	// Get user from JWT
	user_id, err := h.GetUserID(c)
	if err != nil {
		h.responseError(c, http.StatusUnauthorized, err, "user token is not valid")
		return
	}

	// Parsing JSON
	if err := c.ShouldBindJSON(&list); err != nil {
		if err.Error() == "EOF" {
			err = errors.New("request body is empty")
		}
		h.responseError(c, http.StatusBadRequest, err, "json parsing error")
		return
	}

	// Create list
	listId, err := h.service.ToDoList.Create(user_id, list)
	if err != nil {
		h.responseError(c, http.StatusInternalServerError, err, "can't create list")
		return
	}

	// Return created list id
	c.JSON(http.StatusCreated, gin.H{
		"id": listId,
	})
}

func (h *Handler) UpdateList(c *gin.Context) {
	// Get user from JWT
	user_id, err := h.GetUserID(c)
	if err != nil {
		h.responseError(c, http.StatusUnauthorized, err, "user token is not valid")
		return
	}

	// Get list id from url params
	id := c.Param("id")
	list_id, err := strconv.Atoi(id)
	if err != nil {
		h.responseError(c, http.StatusBadRequest, err, fmt.Sprintf("list id %s is invalid", id))
		return
	}

	var list models.UpdateToDoList

	// Parsing JSON
	if err := c.ShouldBindJSON(&list); err != nil {
		if err.Error() == "EOF" {
			err = errors.New("request body is empty")
		}
		h.responseError(c, http.StatusBadRequest, err, "json parsing error")
		return
	}

	err = h.service.ToDoList.UpdateList(user_id, list_id, list)
	if err != nil {
		if strings.Contains(err.Error(), "list not found") {
			h.responseError(c, http.StatusNotFound, err, fmt.Sprintf("list with id %d of user id %d is not found", list_id, user_id))
			return
		}
		h.responseError(c, http.StatusInternalServerError, err, fmt.Sprintf("can't update list id %d of user id %d", list_id, user_id))
		return
	}

	// Return list
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})

}

func (h *Handler) DeleteList(c *gin.Context) {
	// Get user from JWT
	user_id, err := h.GetUserID(c)
	if err != nil {
		h.responseError(c, http.StatusUnauthorized, err, "user token is not valid")
		return
	}

	// Get list id from url params
	id := c.Param("id")
	list_id, err := strconv.Atoi(id)
	if err != nil {
		h.responseError(c, http.StatusBadRequest, err, fmt.Sprintf("list id %s is invalid", id))
		return
	}

	// Delete list
	err = h.service.ToDoList.DeleteList(user_id, list_id)
	if err != nil {
		if strings.Contains(err.Error(), "list not found") {
			h.responseError(c, http.StatusNotFound, err, fmt.Sprintf("list with id %d of user id %d is not found", list_id, user_id))
			return
		}
		h.responseError(c, http.StatusInternalServerError, err, fmt.Sprintf("can't delete list id %d of user id %d", list_id, user_id))
		return
	}

	// Return list
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
