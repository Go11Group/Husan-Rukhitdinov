package inserts

import (
	"encoding/json"
	"fmt"
	crudusers "gorilla/crudUsers"
	"gorilla/model"
	"gorilla/storage/postgres"
	"net/http"

	"github.com/gorilla/mux"
)

type NewUserHttp struct {
	User *crudusers.CrudUsersRepo
}

func NewUserInsert() *crudusers.CrudUsersRepo {
	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}
	userInsert := crudusers.CrudUsersRepo{}
	userInsert.Db = db
	return &userInsert
}

// USER
func CreateUser(w http.ResponseWriter, r *http.Request) {

	userInsert := NewUserInsert()
	user := model.Users{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = userInsert.CreateUsers(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	userInsert := NewUserInsert()
	param := mux.Vars(r)
	id := param["id"]
	user := model.Users{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = userInsert.UpdateUsers(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	err = userInsert.UpdateUsers(user)
	if err != nil {
		_, err = w.Write([]byte("is user  not updated "))
		w.WriteHeader(http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(" user  is  updated "))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	userInsert := NewUserInsert()
	param := mux.Vars(r)
	fmt.Println(param["id"])
	id := param["id"]
	err := userInsert.DeleteUsers(id)

	if err != nil {
		fmt.Println("user is not deleted")
		_, err = w.Write([]byte("Is  not success deleted user"))
		w.WriteHeader(http.StatusBadRequest)
	}
	fmt.Println("user is deleted ")
	_, err = w.Write([]byte("Is   success deleted user"))
	w.WriteHeader(http.StatusOK)

}
