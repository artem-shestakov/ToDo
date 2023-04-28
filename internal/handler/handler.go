package handler

import (
	"github.com/artem-shestakov/to-do/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	service   *service.Service
	logger    *logrus.Logger
	api_token string
}

func NewHadler(service *service.Service, logger *logrus.Logger, api_token string) *Handler {
	return &Handler{
		service:   service,
		logger:    logger,
		api_token: api_token,
	}
}

func (h *Handler) InitRouters() *gin.Engine {
	router := gin.New()
	router.Use(h.Logger())
	router.Use(h.checkAppToken())

	// Auth group
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.SignUp)
		auth.POST("/get-user", h.GetUser)
	}

	// API v1 group
	v1 := router.Group("/api/v1")
	{
		lists := v1.Group("/lists")
		{
			lists.GET("/", h.GetLists)
			lists.GET("/:id", h.GetList)
			lists.POST("/", h.CreateList)
			lists.PUT("/:id", h.UpdateList)
			lists.DELETE("/:id", h.DeleteList)

			tasks := lists.Group(":id/tasks")
			{
				tasks.GET("/", h.GetTasks)
				tasks.GET("/:task_id", h.GetTask)
				tasks.POST("/", h.CreateTask)
				tasks.PUT("/:task_id", h.UpdateTask)
				tasks.DELETE("/:task_id", h.DeleteTask)
			}
		}
	}

	return router
}

func (h *Handler) responseError(c *gin.Context, code int, err error, detail string) {
	// Log
	h.logger.WithFields(logrus.Fields{
		"err": err.Error(),
	}).Errorf(detail)

	// Error response
	c.JSON(code, gin.H{
		"error":  err.Error(),
		"detail": detail,
	})
}
