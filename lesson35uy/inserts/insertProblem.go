package inserts

import (
	"encoding/json"
	"fmt"

	// "my_pro/crudproblems"
	crudproblems "my_pro/crudProblems"
	"my_pro/model"
	"my_pro/storage/postgres"
	"net/http"

	"github.com/gorilla/mux"
)

func ProblemRepo() *crudproblems.CrudProblemsRepo {
	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}
	problem := crudproblems.CrudProblemsRepo{}
	problem.Db = db
	return &problem
}

// PROBLEM
func CreateProblem(w http.ResponseWriter, r *http.Request) {

	var problem model.Problem
	problemRepo := ProblemRepo()

	err := json.NewDecoder(r.Body).Decode(&problem)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		_, err = w.Write([]byte("problem   not created"))
	}
	problemRepo.CreateProblems(problem)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
	}
	_, err = w.Write([]byte("problem    created"))
}

func UpdateProblem(w http.ResponseWriter, r *http.Request) {
	problems := ProblemRepo()
	param := mux.Vars(r)
	id := param["id"]
	var problem model.Problem
	err := json.NewDecoder(r.Body).Decode(&problem)
	if err != nil {
		_, err = w.Write([]byte("problem   not updated"))

		w.WriteHeader(http.StatusInternalServerError)
	}
	err = problems.UpdateProblems(id, problem)
	if err != nil {
		_, err = w.Write([]byte("problem   not updated"))

		w.WriteHeader(http.StatusInternalServerError)
	}
	_, err = w.Write([]byte("problem    updated"))

	w.WriteHeader(http.StatusOK)
}

func DeleteProblem(w http.ResponseWriter, r *http.Request) {

	problems := ProblemRepo()
	param := mux.Vars(r)
	id := param["id"]
	fmt.Println(id)
	err := problems.DeleteProblems(id)
	if err != nil {
		_, err = w.Write([]byte("problem is  not deleted"))
		w.WriteHeader(http.StatusBadRequest)
	}
	_, err = w.Write([]byte("problem    deleted"))

	w.WriteHeader(http.StatusOK)
}

func ReadAllProblem(w http.ResponseWriter, r *http.Request) {
	problemHand := ProblemRepo()
	a, err := problemHand.ReadProblems()
	if err != nil {
		panic(err)
	}
	problems, err := json.Marshal(a)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	_, err = w.Write(problems)
	if err != nil {
		_, err = w.Write([]byte("problem   is getAll "))

		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)

}
