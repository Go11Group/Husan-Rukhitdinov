package main

import (
	"fmt"
	"user_pro/api/handler"
	"user_pro/storage/postgres"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}

	defer db.Close()
	fmt.Println("Connection!!!")

	user := postgres.NewUserRepo(db)

	n := handler.UserRoutes(db, user)
	panic(n.ListenAndServe())
}
