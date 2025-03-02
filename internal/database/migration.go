package database

import "gorm.io/gorm"

// RunMigrations ejecuta las migraciones de la base de datos
func RunMigrations(db *gorm.DB) error {
	return db.AutoMigrate(&Conversation{})
}
