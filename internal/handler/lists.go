package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetLists(c *gin.Context) {
	id, _ := c.Get("userId")

	c.JSON(http.StatusOK, gin.H{
		"user_id": id,
	})
}

func (h *Handler) GetList(c *gin.Context) {

}

func (h *Handler) CreateList(c *gin.Context) {

}

func (h *Handler) UpdateList(c *gin.Context) {

}

func (h *Handler) DeleteList(c *gin.Context) {

}
