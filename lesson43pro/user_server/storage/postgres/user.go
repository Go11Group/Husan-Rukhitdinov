package postgres

import (
	"database/sql"
	"user_pro/models"

	"github.com/google/uuid"
)

type UserRepo struct {
	Db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{Db: db}
}

func (u *UserRepo) Create(user *models.User) error {
	_, err := u.Db.Exec("insert into user (id,name,phone,age) values ($1, $2, $3, $4)",
		uuid.NewString(), user.Name, user.Phone, user.Age)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepo) GetById(id string) (*models.User, error) {
	var user = models.User{Id: id}

	err := u.Db.QueryRow("select name, phone,age from user where id = $1", id).Scan(&user.Name, &user.Phone, &user.Age)
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (u *UserRepo) Update(user *models.User, id string) error {
	_, err := u.Db.Exec("update user set id = $1, name = $2, phone = $3,age = $4", user.Id, user.Name, user.Phone, user.Age)
	if err != nil {
		return err
	}
	return nil
}

func (d *UserRepo) Delete(id string) error {
	_, err := d.Db.Exec("delete from user where id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (c *UserRepo) GetAll(f models.User) ([]models.User, error) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
	)
	query := `select id, name from user `
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
	rows, err := c.Db.Query(query, arr...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err = rows.Scan(&user.Name, &user.Phone, &user.Age)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
