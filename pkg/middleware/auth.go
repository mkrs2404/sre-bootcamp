package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Auth Middleware to authenticate the API Key supplied in the header
func Auth(apiKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := strings.TrimSpace(c.GetHeader("X-API-KEY"))
		if len(apiKey) == 0 {
			err := errors.New("missing API key")
			c.JSON(http.StatusUnauthorized, gin.H{
				"errorMessage": err.Error(),
			})
			c.Abort()
			return
		}

		if apiKey != apiKey {
			err := errors.New("invalid API key")
			c.JSON(http.StatusUnauthorized, gin.H{
				"errorMessage": err.Error(),
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
