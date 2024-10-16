package utils

import (
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

// Logging middleware
func Logger() gin.HandlerFunc {
	file, err := os.OpenFile("trankhacnhu.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to log to file: %v", err)
	}
	logger := log.New(file, "", log.LstdFlags)

	return func(c *gin.Context) {
		start := time.Now()

		// Process request
		c.Next()

		// Log the details after request is processed
		duration := time.Since(start)
		statusCode := c.Writer.Status()
		logger.Printf("| %3d | %13v | %15s | %-7s %#v\n",
			statusCode,
			duration,
			c.ClientIP(),
			c.Request.Method,
			c.Request.URL.Path,
		)
	}
}
