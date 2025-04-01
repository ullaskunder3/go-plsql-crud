package main

import (
	"fmt"

	"github.com/ullaskunder3/go-postgres-crud/db"
)

func main() {
	fmt.Println("🚀 Starting the application...")
	db.InitDB()
}
