package handler

import (
	"database/sql"
	"fmt"
	"net/http"

	"my_pro/storage/postgres"

	"github.com/gorilla/mux"
)

type Handler struct {
	db            *sql.DB
	user          *postgres.CrudUsersRepo
	problem       *postgres.CrudProblemsRepo
	solvedproblem *postgres.CrudSolvedProblemRepo
}

func NewHandler(db *sql.DB, user *postgres.CrudUsersRepo, problem *postgres.CrudProblemsRepo, solved *postgres.CrudSolvedProblemRepo) *Handler {
	return &Handler{db: db, user: user, problem: problem, solvedproblem: solved}
}

func (h *Handler) HandlerAPI() {
	m := mux.NewRouter()
	us := m.PathPrefix("/api/user").Subrouter()
	us.HandleFunc("/salom", func(w http.ResponseWriter, request *http.Request) {
		_, err := w.Write([]byte("salom"))
		if err != nil {
			fmt.Println(err)
		}
	})
    
	us.HandleFunc("/createuser", h.CreateUser).Methods("POST")
	us.HandleFunc("/updateuser/{id}", h.UpdateUser).Methods("PUT")
	us.HandleFunc("/deleteuser/{id}", h.DeleteUser).Methods("DELETE")
	us.HandleFunc("/readuser", h.ReadAllUser).Methods("GET")

	pr := m.PathPrefix("/api/problem").Subrouter()
	pr.HandleFunc("/createproblems", h.CreateProblem).Methods("POST")
	pr.HandleFunc("/updateproblems/{id}", h.UpdateProblem).Methods("PUT")
	pr.HandleFunc("/deleteproblems/{id}", h.DeleteProblem).Methods("DELETE")
	pr.HandleFunc("/readproblems", h.ReadAllProblem).Methods("GET")

	spr := m.PathPrefix("/api/solvedproblem").Subrouter()
	spr.HandleFunc("/createsolved_problems", h.CreateSProblem).Methods("POST")
	spr.HandleFunc("/updatesolved_problems/{id}", h.UpdateSProblem).Methods("PUT")
	spr.HandleFunc("/deletesolved_problems/{id}", h.DeleteSProblem).Methods("DELETE")
	spr.HandleFunc("/readsolved_problems", h.ReadAllSProblem).Methods("GET")

	err := http.ListenAndServe(":8080", m)
	if err != nil {
		fmt.Println(err)
	}
}

