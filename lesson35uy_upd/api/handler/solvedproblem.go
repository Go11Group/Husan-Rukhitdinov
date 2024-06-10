package handler

import (
	"encoding/json"
	"fmt"
	"my_pro/model"
	"my_pro/storage/postgres"
	"net/http"

	"github.com/gorilla/mux"
)

func (h *Handler) NewSolProblemInsert() *postgres.CrudSolvedProblemRepo {
	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}
	problemInsert := postgres.CrudSolvedProblemRepo{}
	problemInsert.Db = db
	return &problemInsert
}

// SOLVED_PROBLEM
func (h *Handler) CreateSProblem(w http.ResponseWriter, r *http.Request) {

	// solvedProblem := NewSolProblemInsert()
	solveProblem := model.SolvedProblem{}

	err := json.NewDecoder(r.Body).Decode(&solveProblem)
	if err != nil {
		// fmt.Println(err)
		w.WriteHeader(http.StatusBadGateway)
	}
	fmt.Println("=========================", solveProblem.ProblemID)
	fmt.Println("+++++++++++++++", solveProblem.UserID)
	err = h.solvedproblem.CreateSolvedProblems(solveProblem)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadGateway)
	}
	_, err = w.Write([]byte("QO'SHILDI!"))
}

func (h *Handler) UpdateSProblem(w http.ResponseWriter, r *http.Request) {
	// solproblemInsert := NewSolProblemInsert()
	param := mux.Vars(r)
	id := param["id"]
	solProblem := model.SolvedProblem{}
	err := json.NewDecoder(r.Body).Decode(&solProblem)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	err = h.solvedproblem.UpdateSolvedProblems(solProblem, id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) DeleteSProblem(w http.ResponseWriter, r *http.Request) {
	// solproblemInsert := NewSolProblemInsert()
	param := mux.Vars(r)
	id := param["id"]
	fmt.Println(id)
	err := h.solvedproblem.DeleteSolvedProblems(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) ReadAllSProblem(w http.ResponseWriter, r *http.Request) {
	// problemSolveHand := NewSolProblemInsert()
	err, a := h.solvedproblem.ReadSolvedProblems()
	if err != nil {
		panic(err)
	}
	fmt.Println(a)
	solve_problems, err := json.Marshal(a)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	_, err = w.Write(solve_problems)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)

}
