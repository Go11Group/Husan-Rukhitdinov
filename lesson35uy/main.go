package main

import (
	"fmt"
	crudproblem "gorilla/crudProblem"
	crudsolvedproblems "gorilla/crudSolvedProblems"
	crudusers "gorilla/crudUsers"
	"gorilla/storage/postgres"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func just(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, this is the root endpoint!"))
}

func just2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, this is the /user/me endpoint!"))
}

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	user := crudusers.NewUsersRepo(db)
	problem := crudproblem.NewProblemsRepo(db)
	solved_problem := crudsolvedproblems.NewSolvedProblemsRepo(db)

	r := mux.NewRouter()

	r.HandleFunc("/readuser", just).Methods(http.MethodGet)

	r2 := r.PathPrefix("/user").Subrouter()
	r2.HandleFunc("/me", just2).Methods(http.MethodGet)

	http.Handle("/", r)
	fmt.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", r))

	fmt.Println("users",user)
	fmt.Println("problems",problem)
	fmt.Println("solved_problems",solved_problem)
}
