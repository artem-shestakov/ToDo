package handler

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) checkAppToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		api_token := c.GetHeader("X-Auth")
		if api_token == "" {
			err := errors.New("app header was not found in request")
			h.responseError(c, http.StatusUnauthorized, err, "Header X-Auth is not found")
			c.Abort()
			return
		}

		if api_token != h.api_token {
			err := errors.New("api token is not valid")
			h.responseError(c, http.StatusUnauthorized, err, "api token is not valid")
			c.Abort()
			return
		}
	}
}

func (h *Handler) GetUserID(c *gin.Context) (int, error) {
	userID := c.Request.Header["X-User"][0]

	if userID == "" {
		return 0, errors.New("user ID if not found in request")
	}
	id, err := strconv.Atoi(userID)
	if err != nil {
		return 0, errors.New("user ID is of invalid type")
	}
	return id, nil
}

func (h *Handler) Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		c.Next()

		latency := time.Since(t)
		status := c.Writer.Status()

		h.logger.WithFields(logrus.Fields{
			"code":    status,
			"path":    c.Request.RequestURI,
			"method":  c.Request.Method,
			"latency": latency,
		}).Info()
	}
}
