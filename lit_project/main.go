package main

import (
	"lit_project/postgres"
	crudproduct "lit_project/crudProducts"
	crudusers "lit_project/crudUsers"
	handler "lit_project/handlers"
	"log"
)

func main() {
	// Ma'lumotlar bazasiga ulanish
	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatalf("Ma'lumotlar bazasiga ulanishda xatolik: %v", err)
	}

	user := crudusers.NewCrudUsersRepo(db)
	product := crudproduct.NewCrudProductRepo(db)

	handlerUser := handler.HandlerUser{User: user, Product: product}

	server := handler.NewHandler(handlerUser)
	server.Addr = ":8080"
	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
