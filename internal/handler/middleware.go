package handler

import (
	"errors"
	"fmt"
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
			responseError(c, http.StatusUnauthorized, fmt.Errorf("authorization header was not found in request"))
			c.Abort()
			return
		}

		// Get token
		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 {
			responseError(c, http.StatusUnauthorized, fmt.Errorf("check authorization Bearer token format"))
			c.Abort()
			return
		}

		// Parse token
		userID, err := h.service.Auth.ParseToken(bearerToken[1])
		if err != nil {
			responseError(c, http.StatusUnauthorized, err)
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

		// Set example variable
		c.Set("example", "12345")

		// before request

		c.Next()

		// after request
		latency := time.Since(t)

		// access the status we are sending
		status := c.Writer.Status()

		h.logger.WithFields(logrus.Fields{
			"code":    status,
			"path":    c.Request.RequestURI,
			"latency": latency,
		}).Info()
	}
}
