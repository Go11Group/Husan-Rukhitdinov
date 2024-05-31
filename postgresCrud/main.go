package main

import (
	"moduls/model"
	"moduls/storage"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	db, err := gorm.Open(postgres.Open("postgres://postgres:3333@localhost:5432/postgres?sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(model.User{})

	// Uncomment to execute specific operations
	// storage.CreateNewUser(db)
	// storage.CreateMultipleUsers(db)
	storage.ReadUserData(db)
	// storage.UpdateUser(db)
	// storage.DeleteUser(db)

	// Create a new user

	// Create multiple users

	// Read user data

	// Update user data

	// Delete user data
}
