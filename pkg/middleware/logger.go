package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		time := time.Now()
		path := c.Request.URL.Path
		verb := c.Request.Method
		

		c.Next()

		var size int
		if c.Writer != nil {
			size = c.Writer.Size()
		}

		fmt.Printf("\n---Request Info------------------------------------------\nTime: %v\nPath: localhost:8080%s\nVerb: %s\nSize: %d\n---------------------------------------------------------\n", time, path, verb, size)
	}
}
