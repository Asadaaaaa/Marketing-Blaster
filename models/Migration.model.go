package models

import (
	"fmt"
)

func RunMigration() {
	err := DB.AutoMigrate(&User{})
	if err != nil {
		panic("Failed to migrate User database")
	}

	errMigrateMailLog := DB.AutoMigrate(&MailLog{})
	if errMigrateMailLog != nil {
		panic("Failed to migrate MailLog database")
	}

	errMigrateAdmin := DB.AutoMigrate(&Admin{})
	if errMigrateAdmin != nil {
		panic("Failed to migrate Admin database")
	}

	fmt.Println("Database Migrated")
}
