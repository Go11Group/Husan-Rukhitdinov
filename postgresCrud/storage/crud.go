package storage

import (
	"fmt"
	"log"
	"moduls/model"

	"gorm.io/gorm"
)

func CreateNewUser(db *gorm.DB) {
	newUser := model.User{FirstName: "Hannah", LastName: "Clark", Email: "hannah.clark@example.com", Password: "password423", Age: 27, Field: "Golang developer", Gender: "Male", IsEmployee: false}
	result := db.Create(&newUser)

	if result.Error != nil {
		log.Fatal(result.Error)
	}
	log.Println("New user ID:", newUser.ID)
}

func CreateMultipleUsers(db *gorm.DB) {
	users := []*model.User{
		{FirstName: "John", LastName: "Doe", Email: "john.doe@example.com", Password: "password123", Age: 28, Field: "Engineering", Gender: "Male", IsEmployee: true},
		{FirstName: "Jane", LastName: "Smith", Email: "jane.smith@example.com", Password: "password123", Age: 34, Field: "Marketing", Gender: "Female", IsEmployee: true},
		{FirstName: "Alice", LastName: "Johnson", Email: "alice.johnson@example.com", Password: "password123", Age: 23, Field: "Design", Gender: "Female", IsEmployee: false},
		{FirstName: "Bob", LastName: "Brown", Email: "bob.brown@example.com", Password: "password123", Age: 45, Field: "Sales", Gender: "Male", IsEmployee: true},
		{FirstName: "Charlie", LastName: "Davis", Email: "charlie.davis@example.com", Password: "password123", Age: 31, Field: "Support", Gender: "Non-binary", IsEmployee: false},
		{FirstName: "Diana", LastName: "Evans", Email: "diana.evans@example.com", Password: "password123", Age: 29, Field: "Human Resources", Gender: "Female", IsEmployee: true},
		{FirstName: "Ethan", LastName: "Garcia", Email: "ethan.garcia@example.com", Password: "password123", Age: 38, Field: "Engineering", Gender: "Male", IsEmployee: true},
		{FirstName: "Fiona", LastName: "Hall", Email: "fiona.hall@example.com", Password: "password123", Age: 27, Field: "Design", Gender: "Female", IsEmployee: false},
		{FirstName: "George", LastName: "Lee", Email: "george.lee@example.com", Password: "password123", Age: 33, Field: "Sales", Gender: "Male", IsEmployee: true},
		{FirstName: "Hannah", LastName: "Miller", Email: "hannah.miller@example.com", Password: "password123", Age: 40, Field: "Support", Gender: "Female", IsEmployee: false},
		{FirstName: "Ian", LastName: "Martinez", Email: "ian.martinez@example.com", Password: "password123", Age: 36, Field: "Human Resources", Gender: "Male", IsEmployee: true},
		{FirstName: "Jenna", LastName: "Nelson", Email: "jenna.nelson@example.com", Password: "password123", Age: 25, Field: "Marketing", Gender: "Female", IsEmployee: true},
		{FirstName: "Kevin", LastName: "Perez", Email: "kevin.perez@example.com", Password: "password123", Age: 44, Field: "Engineering", Gender: "Male", IsEmployee: false},
		{FirstName: "Laura", LastName: "Robinson", Email: "laura.robinson@example.com", Password: "password123", Age: 32, Field: "Design", Gender: "Female", IsEmployee: true},
		{FirstName: "Mike", LastName: "Scott", Email: "mike.scott@example.com", Password: "password123", Age: 28, Field: "Sales", Gender: "Male", IsEmployee: true},
		{FirstName: "Nina", LastName: "Thomas", Email: "nina.thomas@example.com", Password: "password123", Age: 37, Field: "Support", Gender: "Female", IsEmployee: false},
		{FirstName: "Oscar", LastName: "Walker", Email: "oscar.walker@example.com", Password: "password123", Age: 29, Field: "Human Resources", Gender: "Male", IsEmployee: true},
		{FirstName: "Paul", LastName: "Young", Email: "paul.young@example.com", Password: "password123", Age: 41, Field: "Marketing", Gender: "Male", IsEmployee: true},
		{FirstName: "Quinn", LastName: "Allen", Email: "quinn.allen@example.com", Password: "password123", Age: 24, Field: "Engineering", Gender: "Non-binary", IsEmployee: false},
		{FirstName: "Rachel", LastName: "King", Email: "rachel.king@example.com", Password: "password123", Age: 35, Field: "Design", Gender: "Female", IsEmployee: true},
		{FirstName: "Sam", LastName: "Wright", Email: "sam.wright@example.com", Password: "password123", Age: 30, Field: "Sales", Gender: "Male", IsEmployee: true},
		{FirstName: "Tina", LastName: "Harris", Email: "tina.harris@example.com", Password: "password123", Age: 39, Field: "Support", Gender: "Female", IsEmployee: false},
		{FirstName: "Ulysses", LastName: "Clark", Email: "ulysses.clark@example.com", Password: "password123", Age: 42, Field: "Human Resources", Gender: "Male", IsEmployee: true},
		{FirstName: "Ulysses", LastName: "Clark", Email: "ulysses.clark@example.com", Password: "password123", Age: 42, Field: "Human Resources", Gender: "Male", IsEmployee: true},
	}
	for _, user := range users {
		result := db.Create(&user)
		if result.Error != nil {
			log.Fatal(result.Error)
		}
	}
}

func ReadUserData(db *gorm.DB) {
	// var user model.User
	// db.First(&user, 5)
	// fmt.Println(user)

	var persons []model.User
	db.Where("first_name = ?", "Ulysses").Find(&persons)
	fmt.Println(persons)
}

func UpdateUser(db *gorm.DB) {
	var user model.User
	db.First(&user, 5)
	db.Model(&user).Updates(model.User{FirstName: "Oscar", Email: "oscar.walker@example.com"})
	fmt.Println(user)
}

func DeleteUser(db *gorm.DB) {
	var user model.User
	db.First(&user, 5)
	db.Delete(&user)
}
