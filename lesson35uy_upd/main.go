package main

import (
	"fmt"
	"my_pro/api/handler"
	"my_pro/storage/postgres"

	_ "github.com/lib/pq"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		fmt.Println("Error connect")
		return
	}
	defer db.Close()

	userRepo := postgres.NewUsersRepo(db)
	problemRepo := postgres.NewProblemsRepo(db)
	solvedProblemRepo := postgres.NewSolvedProblemsRepo(db)

	h := handler.NewHandler(db, userRepo, problemRepo, solvedProblemRepo)

	h.HandlerAPI()

}
