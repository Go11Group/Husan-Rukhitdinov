package postgres

import (
	"database/sql"
	"fmt"
	"my_pro/model"
)

type CrudUsersRepo struct {
	Db *sql.DB
}

func NewUsersRepo(db *sql.DB) *CrudUsersRepo {
	return &CrudUsersRepo{Db: db}
}

func (c *CrudUsersRepo) CreateUsers(user model.Users) error{
	_, err := c.Db.Exec("insert into users(username,email,password_hash,created_at) values($1,$2,$3,$4)",
		&user.Username, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		fmt.Println(err)
		return  err
	}
	return err
}

func (u *CrudUsersRepo) UpdateUsers(user model.Users,id string) error {
	_, err := u.Db.Exec("update users set username=$1, email=$2, password_hash=$3, created_at=$4 where user_id=$5",
		&user.Username, &user.Email, &user.Password, &user.CreatedAt, &id)
	if err != nil {
		return err
	}
	return err
}

func (d *CrudUsersRepo) DeleteUsers(id string) error{

	_, err := d.Db.Exec("delete from users where user_id=$1", &id)
	if err != nil {
		return err 
	}
	return err
}

func (re *CrudUsersRepo) ReadUsers() ([]model.Users, error) {
	row, err := re.Db.Query("SELECT user_id, username, email, password_hash, created_at FROM users")
	if err != nil {
		return nil, err
	}
	defer row.Close()

	var users []model.Users
	for row.Next() {
		var user model.Users
		err = row.Scan(&user.UserID, &user.Username, &user.Email, &user.Password, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err = row.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

