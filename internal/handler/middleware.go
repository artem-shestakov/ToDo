package handler

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
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
