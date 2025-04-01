package main

import (
	"fmt"

	"github.com/ullaskunder3/go-postgres-crud/db"
	"github.com/ullaskunder3/go-postgres-crud/routes"
)

func main() {
	fmt.Println("🚀 Starting the application...")
	db.InitDB()
	db.MigrateDB()

	// Start Gin server
	r := routes.SetupRouter()
	fmt.Println("🌍 Server running on http://localhost:8080")
	r.Run(":8080")
}
