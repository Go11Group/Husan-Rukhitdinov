package handler

import (
	crudproblem "gorilla/crudProblem"
	crudSolvedProblems "gorilla/crudSolvedProblems"
	crudusers "gorilla/crudUsers"
	"gorilla/inserts"
	"gorilla/storage/postgres"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	User           *crudusers.CrudUsersRepo
	Problem        *crudproblem.CrudProblemsRepo
	Solved_problem *crudSolvedProblems.CrudSolvedProblemRepo
}

type HandlerUsers struct {
	User *crudusers.CrudUsersRepo
}

type HandlerProblems struct {
	Problem *crudproblem.CrudProblemsRepo
}

type HandlerSolvedProblems struct {
	SolvedPro *crudSolvedProblems.CrudSolvedProblemRepo
}

func NewHandler(user *crudusers.CrudUsersRepo) *http.Server {
	m := mux.NewRouter()
	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	use:=crudusers.NewUsersRepo(db)
	all:=inserts.NewHttpMod(use)
	
	// mux := http.NewServeMux()
	// handler := HandlerUsers{User: user}
	m.HandleFunc("/createuser", all.User.CreateUsers).Methods("POST")
	m.HandleFunc("/updateuser", all.User.UpdateUsers).Methods("PUT")
	m.HandleFunc("/deleteuser", all.User.DeleteUsers).Methods("DELETE")
	m.HandleFunc("/readuser", all.User.ReadUsers).Methods("GET")

	return &http.Server{Handler: m}
}

func NewHandlerPro(problem *crudproblem.CrudProblemsRepo) *http.Server {
	m := mux.NewRouter()
	handler := HandlerProblems{Problem: problem}
	m.HandleFunc("/createproblems", handler.Problem.CreateProblems).Methods("POST")
	m.HandleFunc("/updateproblems", handler.Problem.UpdateProblems).Methods("PUT")
	m.HandleFunc("/dalateproblems", handler.Problem.DeleteProblems).Methods("DELETE")
	m.HandleFunc("/readproblems", handler.Problem.ReadProblems).Methods("GET")

	return &http.Server{Handler: m}
}

func NewHandlerSolvedPro(solved_problem *crudSolvedProblems.CrudSolvedProblemRepo) *http.Server {
	m := mux.NewRouter()
	handler := HandlerSolvedProblems{SolvedPro: solved_problem}
	m.HandleFunc("/createsolved_problems", handler.SolvedPro.CreateSolvedProblems).Methods("POST")
	m.HandleFunc("/updatesolved_problems", handler.SolvedPro.UpdateSolvedProblems).Methods("PUT")
	m.HandleFunc("/dalatesolved_problems", handler.SolvedPro.DeleteSolvedProblems).Methods("DELETE")
	m.HandleFunc("/readsolved_problems", handler.SolvedPro.ReadSolvedProblems).Methods("GET")

	return &http.Server{Handler: m}
}
