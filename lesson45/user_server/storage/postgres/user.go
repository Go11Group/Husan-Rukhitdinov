package postgres

import (
	"database/sql"
	"time"
	pb "github.com/Go11Group/Husan-Rukhitdinov/lesson44/user_server/genproto/user_service"
)

type UserRepo struct {
	Db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{Db: db}
}

func (u *UserRepo) Create(user *pb.User) error {
	_, err := u.Db.Exec("insert into user (name,phone,age,created_at) values ($1, $2, $3, $4)",
		user.Name, user.Phone, user.Age, time.Now())
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepo) GetById(id string) (*pb.User, error) {
	var user = pb.User{Id: id}

	err := u.Db.QueryRow("select name, phone,age from user where id = $1", id).Scan(&user.Name, &user.Phone, &user.Age)
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (u *UserRepo) Update(user *pb.User, id string) error {
	_, err := u.Db.Exec("update user set name = $1, phone = $2,age = $3,update_at = $4 where id = $5", user.Name, user.Phone, &user.Age, time.Now(), id)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepo) Delete(id string) error {
	_, err := u.Db.Exec("update user set deleted_at = $1 where id = $2", 1, id)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepo) GetAll(f *pb.User) ([]pb.User, error) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
	)
	query := `select id, name,phone,age created_at, update_at,delete_at from user where deleted_at = 0 `
	filter := ` where true `

	if len(f.Name) > 0 {
		params["name"] = f.Name
		filter += "and name = :name "
	}
	if len(f.Phone) > 0 {
		params["phone"] = f.Phone
		filter += "and phone = :phone "
	}
	if f.Age > 0 {
		params["age"] = f.Age
		filter += "and age = :age "
	}

	query = query + filter

	query, arr = ReplaceQueryParam(query, params)
	rows, err := u.Db.Query(query, arr...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []pb.User
	for rows.Next() {
		var user pb.User
		err = rows.Scan(&user.Name, &user.Phone, &user.Age, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

