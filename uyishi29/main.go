package main

import (
	"database/sql"
	"fmt"
	"go_run/model"
	"go_run/storage/postgres"
	"log"

	_ "github.com/lib/pq"
)

// ConnectDB PostgreSQL bilan bog'lanish
func ConnectDB() (*sql.DB, error) {
	connStr := "user=postgres password=3333 dbname=homework sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func main() {
	db, err := ConnectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	userRepo := postgres.NewInfoRepo(db)
	productRepo := postgres.NewInfoRepo(db)

	// CREATE USER
	newUser := model.Users{Username: "Alice", Email: "alice.johnson@example.com", Password: "password123"}
	err = userRepo.CreateNewUser(newUser)
	if err != nil {
		log.Fatalf("Error creating user: %v\n", err)
	}

	// CREATE PRODUCT
	newProduct := model.Product{Name: "MacBook Pro", Discription: "High-performance laptop for professionals", Price: 1999.99, StockQuantity: 50}
	err = productRepo.CreateNewProduct(newUser.Id, newProduct)
	if err != nil {
		log.Fatalf("Error creating product: %v\n", err)
	}

	// READ USERS
	users, err := userRepo.ReadAllUsers()
	if err != nil {
		log.Fatal(err)
	}

	for _, user := range users {
		fmt.Printf("ID: %d, Username: %s, Email: %s\n", user.Id, user.Username, user.Email)
	}

	// READ PRODUCTS
	products, err := productRepo.ReadAllProducts()
	if err != nil {
		log.Fatal(err)
	}

	for _, product := range products {
		fmt.Printf("ID: %d, Name: %s, Description: %s, Price: %v, StockQuantity: %d\n", product.Id, product.Name, product.Discription, product.Price, product.StockQuantity)
	}

	// UPDATE USER
	newUser.Username = "updateduser"
	newUser.Email = "updateduser@example.com"
	err = userRepo.UpdateUser(newUser)
	if err != nil {
		log.Fatalf("Error updating user: %v", err)
	}
	fmt.Println("User updated successfully")

	// UPDATE PRODUCT
	existingProduct := model.Product{
		Id:             newProduct.Id,
		Name:           "Updated Product Name",
		Discription:    "Updated Product Description",
		Price:          29.99,
		StockQuantity:  100,
	}
	err = productRepo.UpdateProduct(existingProduct)
	if err != nil {
		log.Fatalf("Error updating product: %v", err)
	}
	fmt.Println("Product updated successfully")

	// DELETE USER
	err = userRepo.DeleteUser(newUser.Id)
	if err != nil {
		log.Fatalf("Failed to delete user: %v", err)
	}
	fmt.Println("User deleted")

	// DELETE PRODUCT
	err = productRepo.DeleteProduct(newProduct.Id)
	if err != nil {
		log.Fatalf("Failed to delete product: %v", err)
	}
	fmt.Println("Product deleted")
}
