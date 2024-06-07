package crudusers

import (
	"lit_project/inserts"
	"database/sql"
	"encoding/json"
	"lit_project/model"
	"net/http"
)

type CrudUsersRepo struct {
	db *sql.DB
}

func NewCrudUsersRepo(db *sql.DB) *CrudUsersRepo {
	return &CrudUsersRepo{db: db}
}

func (i *CrudUsersRepo) InsertUsers(w http.ResponseWriter, r *http.Request) {
	user, err := Insert_test.InsertUser(w, r)
	if err != nil {
		panic(err)
	}

	_, err = i.db.Exec("insert into users(id,username,email,password) values($1,$2,$3,$4)",
		&user.Id, &user.Firstname, &user.Lastname, &user.Age)
	if err != nil {
		panic(err)
	}
}

func (u *CrudUsersRepo) UpdateUsers(w http.ResponseWriter, r *http.Request) {
	up, err := Insert_test.UpdateUser(w, r)
	if err != nil {
		panic(err)
	}

	_, err = u.db.Exec("update users set username=$1, email=$2,pasword=$3 where id=$4",
		&up.Firstname, &up.Lastname, &up.Age, &up.Id)

	if err != nil {
		panic(err)
	}
}

func (d *CrudUsersRepo) DeleteUsers(w http.ResponseWriter, r *http.Request) {
	del, err := Insert_test.DeleteUser(w, r)
	if err != nil {
		panic(err)
	}

	_, err = d.db.Exec("delete from users where id=$1", &del.Id)
	if err != nil {
		panic(err)
	}
}

func (s *CrudUsersRepo) ReadUsers(w http.ResponseWriter, r *http.Request) {
	row, err := s.db.Query("select * from users")
	if err != nil {
		http.Error(w, "Unable to scan the row", http.StatusInternalServerError)
		return
	}
	defer row.Close()

	var users []model.Users
	for row.Next() {
		var user model.Users
		err = row.Scan(&user.Id, &user.Firstname, &user.Lastname, &user.Age)
		if err != nil {
			http.Error(w, "Unable to scan the row", http.StatusInternalServerError)
			return
		}

		users = append(users, user)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		panic(err)
	}
}
