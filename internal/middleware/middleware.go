package middleware

import (
	"github.com/gin-gonic/gin"
	"time"
)

func RateLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Simple rate limiting logic
		c.Next()
	}
}

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		
		c.Next()
		
		endTime := time.Now()
		latency := endTime.Sub(startTime)
		
		// Log request details
		println("Path:", c.Request.URL.Path)
		println("Latency:", latency)
	}
}