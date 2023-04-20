package handler

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/artem-shestakov/to-do/internal/models"
	"github.com/gin-gonic/gin"
)

type loginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
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

func (h *Handler) Login(c *gin.Context) {
	var login loginInput
	if err := c.ShouldBindJSON(&login); err != nil {
		h.responseError(c, http.StatusBadRequest, err, "JSON parsing error")
		return
	}

	token, err := h.service.Auth.GenerateToken(login.Email, login.Password)
	if err != nil && err == sql.ErrNoRows {
		err = fmt.Errorf("user %s not found or email/password incorrect", login.Email)
		h.responseError(c, http.StatusUnauthorized, err, fmt.Sprintf("Can't authoririze user: %s", login.Email))
		return
	} else if err != nil {
		h.responseError(c, http.StatusUnauthorized, err, fmt.Sprintf("Can't authoririze user: %s", login.Email))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
