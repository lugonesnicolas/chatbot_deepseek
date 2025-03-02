package handlers

import (
	"chatbot-api/internal/database"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// HistoryHandler maneja la recuperación del historial de conversaciones
func HistoryHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var conversations []database.Conversation
		db.Find(&conversations)
		c.JSON(http.StatusOK, conversations)
	}
}
