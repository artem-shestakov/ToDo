package handler

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/artem-shestakov/to-do/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type loginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) SignUp(c *gin.Context) {
	var json models.User
	if err := c.ShouldBindJSON(&json); err != nil {
		responseError(c, http.StatusBadRequest, err)
		h.logger.WithFields(logrus.Fields{
			"err": err.Error(),
		}).Errorf("JSON parsing error")
		return
	}

	id, err := h.service.Auth.CreateUser(json)
	if err != nil {
		responseError(c, http.StatusInternalServerError, err)
		h.logger.WithFields(logrus.Fields{
			"err": err.Error(),
		}).Errorf("User is not created")
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"id": id,
	})

}

func (h *Handler) Login(c *gin.Context) {
	var login loginInput
	if err := c.ShouldBindJSON(&login); err != nil {
		responseError(c, http.StatusBadRequest, err)
		h.logger.WithFields(logrus.Fields{
			"err": err.Error(),
		}).Errorf("JSON parsing error")
		return
	}

	token, err := h.service.Auth.GenerateToken(login.Email, login.Password)
	if err != nil && err == sql.ErrNoRows {
		err = fmt.Errorf("user %s not found or email/password incorrect", login.Email)
		responseError(c, http.StatusUnauthorized, err)
		h.logger.WithFields(logrus.Fields{
			"err": err.Error(),
		}).Errorf("Can't authoririze user: %s", login.Email)
		return
	} else if err != nil {
		responseError(c, http.StatusUnauthorized, err)
		h.logger.WithFields(logrus.Fields{
			"err": err.Error(),
		}).Errorf("Can't authoririze user: %s", login.Email)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
