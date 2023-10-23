package migrations

import (
	"Marketing-Blaster/models"
	"fmt"
)

func RunMigration() {
	err := models.DB.AutoMigrate(&models.User{})
	if err != nil {
		panic("Failed to migrate database")
	}

	fmt.Println("Database Migrated")
}
