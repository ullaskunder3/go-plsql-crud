package repositories

import (
	"github.com/ullaskunder3/go-postgres-crud/db"
	"github.com/ullaskunder3/go-postgres-crud/models"
)

// CreateUser inserts a new user into the database
func CreateUser(user *models.User) error {
	result := db.DB.Create(user)
	return result.Error
}

// GetUserByID fetches a user by ID
func GetUserByID(id uint) (*models.User, error) {
	var user models.User
	result := db.DB.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// GetAllUsers retrieves all users
func GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := db.DB.Find(&users)
	return users, result.Error
}

// UpdateUser updates an existing user's details
func UpdateUser(user *models.User) error {
	result := db.DB.Save(user)
	return result.Error
}

// DeleteUser removes a user by ID
func DeleteUser(id uint) error {
	result := db.DB.Delete(&models.User{}, id)
	return result.Error
}
