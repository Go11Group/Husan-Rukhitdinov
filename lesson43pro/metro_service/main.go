package main

import (
	"fmt"
	"pro/api"
	"pro/storage/postgres"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}

	defer db.Close()
	fmt.Println("Connection!!!")

	// server := api.Routes(db)
	// err = server.ListenAndServe() // 80
	// if err != nil {
	// 	panic(err)
	// }
	card := postgres.NewCardRepo(db)
	station := postgres.NewStationRepo(db)
	terminal := postgres.NewTerminalRepo(db)
	transaction := postgres.NewTransactionRepo(db)

	k := api.Routes(db, station, card, terminal, transaction)
	panic(k.ListenAndServe())
}
