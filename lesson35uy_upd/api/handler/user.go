package handler

import (
	"encoding/json"
	"fmt"
	"my_pro/model"
	"my_pro/storage/postgres"
	"net/http"

	"github.com/gorilla/mux"
)

func (h *Handler) NewUserInsert() *postgres.CrudUsersRepo {
	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}
	userInsert := postgres.CrudUsersRepo{}
	userInsert.Db = db
	return &userInsert
}

// USER
func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {

	user := model.Users{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.user.CreateUsers(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {

	param := mux.Vars(r)
	id := param["id"]
	user := model.Users{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.user.UpdateUsers(user, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	err = h.user.UpdateUsers(user, id)
	if err != nil {
		_, err = w.Write([]byte("is user  not updated "))
		w.WriteHeader(http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(" user  is  updated "))
}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {

	param := mux.Vars(r)
	fmt.Println(param["id"])
	id := param["id"]
	err := h.user.DeleteUsers(id)

	if err != nil {
		fmt.Println("user is not deleted")
		_, err = w.Write([]byte("Is  not success deleted user"))
		w.WriteHeader(http.StatusBadRequest)
	}
	fmt.Println("user is deleted ")
	_, err = w.Write([]byte("Is   success deleted user"))
	w.WriteHeader(http.StatusOK)

}

func (h *Handler) ReadAllUser(w http.ResponseWriter, r *http.Request) {
	a, err := h.user.ReadUsers()
	if err != nil {
		panic(err)
	}
	fmt.Println("a========", a)
	users, err := json.Marshal(a)
	fmt.Println("user   =====", users)
	if err != nil {
		fmt.Println("Marshal errr-r---")
		w.WriteHeader(http.StatusBadRequest)
	}
	_, err = w.Write(users)
	if err != nil {
		_, err = w.Write([]byte("is user  not getAll "))
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)

}
