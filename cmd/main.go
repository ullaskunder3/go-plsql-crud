package main

import (
	"fmt"

	"github.com/ullaskunder3/go-postgres-crud/db"
	"github.com/ullaskunder3/go-postgres-crud/models"
	"github.com/ullaskunder3/go-postgres-crud/repositories"
)

func main() {
	fmt.Println("🚀 Starting the application...")
	db.InitDB()
	db.MigrateDB()

	// Test CRUD Operations
	testCRUD()
}

func testCRUD() {
	// Create a new user
	newUser := models.User{Name: "John Doe", Email: "john@example.com"}
	err := repositories.CreateUser(&newUser)
	if err != nil {
		fmt.Println("❌ Error creating user:", err)
	} else {
		fmt.Println("✅ User created successfully:", newUser)
	}

	// Get user by ID
	user, err := repositories.GetUserByID(newUser.ID)
	if err != nil {
		fmt.Println("❌ Error fetching user:", err)
	} else {
		fmt.Println("✅ User fetched successfully:", user)
	}

	// Update user
	user.Name = "John Updated"
	err = repositories.UpdateUser(user)
	if err != nil {
		fmt.Println("❌ Error updating user:", err)
	} else {
		fmt.Println("✅ User updated successfully:", user)
	}

	// Get all users
	users, err := repositories.GetAllUsers()
	if err != nil {
		fmt.Println("❌ Error fetching all users:", err)
	} else {
		fmt.Println("✅ All users:", users)
	}

	// Delete user
	err = repositories.DeleteUser(user.ID)
	if err != nil {
		fmt.Println("❌ Error deleting user:", err)
	} else {
		fmt.Println("✅ User deleted successfully")
	}
}
