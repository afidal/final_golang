package middleware

import (
	"github.com/gin-gonic/gin"
	"errors"
	"os"
	"tp_final/pkg/web"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("TOKEN")
		if token == "" {
			web.Failure(c, 401, errors.New("Token not found"))
			c.Abort()
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(c, 401, errors.New("Invalid token"))
			c.Abort()
			return
		}
		c.Next()
	}
}