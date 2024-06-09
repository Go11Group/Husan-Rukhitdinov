package inserts

import (
	"encoding/json"
	"fmt"
	crudsolvedproblems "gorilla/crudSolvedProblems"
	"gorilla/model"
	"gorilla/storage/postgres"
	"net/http"
)

type NewSolvedProblemHttp struct {
	SolProblem *crudsolvedproblems.CrudSolvedProblemRepo
}

func NewSolProblemInsert() *crudsolvedproblems.CrudSolvedProblemRepo {
	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}
	problemInsert := crudsolvedproblems.CrudSolvedProblemRepo{}
	problemInsert.Db = db
	return &problemInsert
}

// SOLVED_PROBLEM
func CreateSProblem(w http.ResponseWriter, r *http.Request) {

	solvedProblem := NewSolProblemInsert()
	solveProblem := model.SolvedProblem{}

	err := json.NewDecoder(r.Body).Decode(&solvedProblem)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadGateway)
	}
	fmt.Println(solvedProblem)
	err = solvedProblem.CreateSolvedProblem(solveProblem)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadGateway)
	}
}

func (us *NewHttp) UpdateSProblem(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "It's not PUT request!", http.StatusMethodNotAllowed)
		return
	}
	var solvedProblem model.SolvedProblem
	err := json.NewDecoder(r.Body).Decode(&solvedProblem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = us.User.UpdateSolvedProblem(solvedProblem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(solvedProblem)
}

func (us *NewHttp) DeleteSProblem(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "It's not DELETE request!", http.StatusMethodNotAllowed)
		return
	}
	var solvedProblem model.SolvedProblem
	err := json.NewDecoder(r.Body).Decode(&solvedProblem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = us.User.DeleteSolvedProblem(solvedProblem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(solvedProblem)
}
