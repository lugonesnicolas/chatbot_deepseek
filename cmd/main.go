package main

import (
	"chatbot-api/internal/database"
	"chatbot-api/internal/handlers"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializar base de datos
	db, err := database.InitDB()
	if err != nil {
		panic(fmt.Sprintf("Error al conectar con la base de datos: %v", err))
	}

	r := gin.Default()

	// Habilitar CORS para permitir que el frontend consuma la API
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Si es una solicitud OPTIONS, responder con 200 OK
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	})

	// Rutas de la API
	r.POST("/chat", handlers.ChatbotHandler(db))
	r.GET("/history", handlers.HistoryHandler(db))

	// Iniciar el servidor
	r.Run("0.0.0.0:8080")
}
