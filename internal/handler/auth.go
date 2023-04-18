package handler

import (
	"net/http"

	"github.com/artem-shestakov/to-do/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

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

}
