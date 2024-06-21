package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"projects/api"
	"projects/storage/postgres"
)

func main() {
	db, err := postgres.ConnectionDb()
	if err != nil {
		panic(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)
	fmt.Println("Connect!!!")
	user := postgres.NewUserRepository(db)
	router := api.RooterApi(user)
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		return
	}
	fmt.Println("Listennig :8080")

}
