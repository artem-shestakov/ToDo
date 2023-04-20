package handler

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get Authorization token
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			err := errors.New("authorization header was not found in request")
			h.responseError(c, http.StatusUnauthorized, err, "Header Authorization is not found")
			c.Abort()
			return
		}

		// Get token
		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 {
			err := errors.New("incorrect Bearer token format")
			h.responseError(c, http.StatusUnauthorized, err, "Can't get Bearer token from header")
			c.Abort()
			return
		}

		// Parse token
		userID, err := h.service.Auth.ParseToken(bearerToken[1])
		if err != nil {
			h.responseError(c, http.StatusUnauthorized, err, "Token parsing error")
			c.Abort()
			return
		}
		// Set context var
		c.Set("userId", userID)
	}
}

func (h *Handler) GetUserID(c *gin.Context) (int, error) {
	userID, _ := c.Get("userId")
	if userID == "" {
		return 0, errors.New("user ID if not found in request")
	}
	id, ok := userID.(int)
	if !ok {
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
