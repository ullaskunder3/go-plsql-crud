package db

import (
	"fmt"

	"github.com/ullaskunder3/go-postgres-crud/models"
)

// MigrateDB runs database migrations
func MigrateDB() {
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		fmt.Println("❌ Failed to migrate database:", err)
	} else {
		fmt.Println("✅ Database migrated successfully!")
	}
}
