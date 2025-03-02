package handlers

import (
	"chatbot-api/internal/database"
	"chatbot-api/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ChatbotHandler maneja la interacción con el chatbot
func ChatbotHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			Message string `json:"message"`
		}

		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		response, err := services.QueryOllama(req.Message)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get response"})
			return
		}

		// Guardar la conversación en la base de datos
		db.Create(&database.Conversation{Message: req.Message, Response: response})

		c.JSON(http.StatusOK, gin.H{"response": response})
	}
}
