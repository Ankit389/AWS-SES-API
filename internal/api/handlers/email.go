package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type EmailRequest struct {
	Source      string   `json:"source" binding:"required,email"`
	Destination []string `json:"destination" binding:"required,dive,email"`
	Message     struct {
		Subject string `json:"subject" binding:"required"`
		Body    string `json:"body" binding:"required"`
	} `json:"message" binding:"required"`
}

func SendEmail(c *gin.Context) {
	var req EmailRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Mock successful response
	c.JSON(http.StatusOK, gin.H{
		"MessageId": "mock-message-id-123",
		"Status":    "Success",
	})
}

func SendRawEmail(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"MessageId": "mock-raw-message-id-123",
		"Status":    "Success",
	})
}

func SendTemplatedEmail(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"MessageId": "mock-template-message-id-123",
		"Status":    "Success",
	})
}

func GetSendQuota(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Max24HourSend":   200.0,
		"MaxSendRate":     1.0,
		"SentLast24Hours": 0.0,
	})
}

func GetSendStatistics(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"SendDataPoints": []gin.H{
			{
				"Timestamp":         "2024-01-20T00:00:00Z",
				"DeliveryAttempts": 0,
				"Bounces":          0,
				"Complaints":       0,
				"Rejects":         0,
			},
		},
	})
}