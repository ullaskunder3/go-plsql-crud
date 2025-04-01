package models

import "gorm.io/gorm"

// User represents a user in the system
type User struct {
	gorm.Model
	Name  string `gorm:"type:varchar(100);not null"`
	Email string `gorm:"type:varchar(100);uniqueIndex;not null"`
}
