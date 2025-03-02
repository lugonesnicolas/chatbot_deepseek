package database

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB inicializa la base de datos SQLite
func InitDB() (*gorm.DB, error) {
	var err error
	DB, err = gorm.Open(sqlite.Open("chatbot.db"), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error al conectar con la base de datos: %w", err)
	}

	// Migrar modelos
	DB.AutoMigrate(&Conversation{})
	fmt.Println("Base de datos SQLite inicializada correctamente")
	return DB, nil
}

// Conversation representa el modelo de la base de datos
type Conversation struct {
	ID       uint   `gorm:"primaryKey"`
	Message  string `json:"message"`
	Response string `json:"response"`
}
