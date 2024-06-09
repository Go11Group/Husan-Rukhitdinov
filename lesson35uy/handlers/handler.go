package handler

import (
	"fmt"
	"my_pro/inserts"
	"net/http"

	"github.com/gorilla/mux"
)

func HandlerAPI() {
	m := mux.NewRouter()
	us := m.PathPrefix("/api/user").Subrouter()
	us.HandleFunc("/salom", func(w http.ResponseWriter, request *http.Request) {
		_, err := w.Write([]byte("salom"))
		fmt.Println(err)
	})

	us.HandleFunc("/createuser", inserts.CreateUser).Methods("POST")
	us.HandleFunc("/updateuser/{id}", inserts.UpdateUser).Methods("PUT")
	us.HandleFunc("/deleteuser/{id}", inserts.DeleteUser).Methods("DELETE")
	us.HandleFunc("/readuser", inserts.ReadAllUser).Methods("GET")

	pr := m.PathPrefix("/api/problem").Subrouter()
	pr.HandleFunc("/createproblems", inserts.CreateProblem).Methods("POST")
	pr.HandleFunc("/updateproblems/{id}", inserts.UpdateProblem).Methods("PUT")
	pr.HandleFunc("/deleteproblems/{id}", inserts.DeleteProblem).Methods("DELETE")
	pr.HandleFunc("/readproblems", inserts.ReadAllProblem).Methods("GET")

	spr := m.PathPrefix("/api/solvedproblem").Subrouter()
	spr.HandleFunc("/createsolved_problems", inserts.CreateSProblem).Methods("POST")
	spr.HandleFunc("/updatesolved_problems/{id}", inserts.UpdateSProblem).Methods("PUT")
	spr.HandleFunc("/deletesolved_problems/{id}", inserts.DeleteSProblem).Methods("DELETE")
	spr.HandleFunc("/readsolved_problems", inserts.ReadAllSProblem).Methods("GET")
	err := http.ListenAndServe(":8080", m)
	if err != nil {
		fmt.Println(err)
	}
}
