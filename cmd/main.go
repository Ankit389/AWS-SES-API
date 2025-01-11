// package main

// import (
// 	"log"
// 	"github.com/gin-gonic/gin"
// 	"aws-ses-mock/internal/api/handlers"
// 	"aws-ses-mock/internal/middleware"
// )

// func main() {
// 	r := gin.Default()
	
// 	// CORS Configuration
// 	r.Use(func(c *gin.Context) {
// 		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
// 		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
// 		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, Authorization")
		
// 		if c.Request.Method == "OPTIONS" {
// 			c.AbortWithStatus(204)
// 			return
// 		}
		
// 		c.Next()
// 	})
	
// 	// Middleware
// 	r.Use(middleware.RateLimiter())
// 	r.Use(middleware.RequestLogger())

// 	// API Routes
// 	v1 := r.Group("/v1")
// 	{
// 		v1.POST("/email/send", handlers.SendEmail)
// 		v1.POST("/email/send-raw", handlers.SendRawEmail)
// 		v1.POST("/email/send-templated", handlers.SendTemplatedEmail)
// 		v1.GET("/email/quota", handlers.GetSendQuota)
// 		v1.GET("/email/statistics", handlers.GetSendStatistics)
// 	}

// 	// Start server
// 	if err := r.Run(":8080"); err != nil {
// 		log.Fatal("Failed to start server:", err)
// 	}
// }



package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"aws-ses-mock/internal/api/handlers"
	"aws-ses-mock/internal/middleware"
)

func main() {
	r := gin.Default()
	
	// CORS Configuration
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, Authorization")
		
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		
		c.Next()
	})
	
	// Middleware
	r.Use(middleware.RateLimiter())
	r.Use(middleware.RequestLogger())

	// API Routes
	v1 := r.Group("/v1")
	{
		v1.POST("/email/send", handlers.SendEmail)
		v1.POST("/email/send-raw", handlers.SendRawEmail)
		v1.POST("/email/send-templated", handlers.SendTemplatedEmail)
		v1.GET("/email/quota", handlers.GetSendQuota)
		v1.GET("/email/statistics", handlers.GetSendStatistics)
	}

	// Start server on port 3000
	if err := r.Run(":3000"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}