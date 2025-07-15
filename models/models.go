package models

import (
	"log"

	"gorm.io/gorm"
)

// AutoMigrate runs database migrations for all models
func AutoMigrate(db *gorm.DB) error {
	log.Println("Running database migrations...")
	
	err := db.AutoMigrate(
		&User{},
	)
	
	if err != nil {
		log.Printf("Migration failed: %v", err)
		return err
	}
	
	log.Println("Database migrations completed successfully")
	return nil
} 