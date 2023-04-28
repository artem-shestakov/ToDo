package handler

import (
	"net/http"

	"github.com/artem-shestakov/to-do/internal/models"
	"github.com/gin-gonic/gin"
)

type userId struct {
	UserId int `json:"user_id"`
}

func (h *Handler) SignUp(c *gin.Context) {
	var json models.User
	if err := c.ShouldBindJSON(&json); err != nil {
		h.responseError(c, http.StatusBadRequest, err, "JSON parsing error")

		return
	}

	id, err := h.service.Auth.CreateUser(json)
	if err != nil {
		h.responseError(c, http.StatusInternalServerError, err, "User is not created")
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"id": id,
	})

}

func (h *Handler) GetUser(c *gin.Context) {
	var json userId
	if err := c.ShouldBindJSON(&json); err != nil {
		h.responseError(c, http.StatusBadRequest, err, "JSON parsing error")

		return
	}

	user, err := h.service.Auth.GetUser(json.UserId)
	if err != nil {
		h.responseError(c, http.StatusNotFound, err, "user is not found")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})

}
