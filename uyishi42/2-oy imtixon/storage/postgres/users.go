package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"project/models"
	"time"

	"github.com/google/uuid"
)

type User struct {
	db *sql.DB
}

func NewUser(db *sql.DB) *User {
	return &User{db}
}

func (u *User) CreateUser(user *models.Users) error {
	birthday, err := time.Parse("2006-01-02", user.Birthday)
	if err != nil {
		fmt.Println("hlooo", err)
		return errors.New("failed to parse birthday")
	}

	user.ID = uuid.NewString()
	_, err = u.db.Exec("insert into users (user_id,name,email,birthday,password, created_at) VALUES ($1,$2,$3,$4,$5,$6)",
		user.ID, user.Name, user.Email, birthday, user.Password, time.Now())
	if err != nil {
		return err
	}
	return nil
}
func (u *User) ReadUser(id string) (*models.Users, error) {
	row := u.db.QueryRow("select * from users where user_id = $1", id)

	fmt.Println(id)
	var user models.Users
	err := row.Scan(
		&user.ID, &user.Name, &user.Email, &user.Birthday, &user.Password, &user.CreatedAt, &user.UpdateAt, &user.DeleteAt)
	if err != nil {

		return nil, err
	}
	return &user, nil
}

func (u *User) UpdateUser(user *models.Users,id string) error {
	birthday, err := time.Parse("2006-01-02", user.Birthday)
	if err != nil {
		fmt.Println("dd.....dd",err)
		return errors.New("failed to parse birthday")
	}
	_, err = u.db.Exec("update users set name = $1, email = $2, birthday = $3, password = $4 where user_id = $5",
		user.Name, user.Email, birthday, user.Password, id)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) DeleteUser(id string) error {
	_, err := u.db.Exec("delete from users where user_id = $1", id)
	if err != nil {
		fmt.Println("dsds--------",err)
		return err
	}
	return nil
}

func (u *User) FilterAllUser(f models.UserF) ([]models.UserF, error) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
	)
	query := `select user_id,name,email,birthday,password
		from users `
	filter := ` where true`

	if len(f.ID) > 0 {
		params["user_id"] = f.ID
		filter += " and user_id = :user_id "
	}
	if len(f.Name) > 0 {
		params["name"] = f.Name
		filter += "and name = :name "
	}

	if len(f.Email) > 0 {
		params["email"] = f.Email
		filter += " and email = :email "
	}

	if len(f.Birthday) > 0 {
		params["birthday"] = f.Birthday
		filter += " and birthday = :birthday "
	}

	if len(f.Password) > 0 {
		params["password"] = f.Password
		filter += " and password = :password "
	}

	query = query + filter

	query, arr = ReplaceQueryParam(query, params)
	rows, err := u.db.Query(query, arr...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.UserF
	for rows.Next() {
		var user models.UserF
		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Birthday, &user.Password)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}
	return users, nil
}
// 1-TASK
func (u *User) GetCourseByUser(userId string) (string, []models.Courses, error) {
	var courses string
	err := u.db.QueryRow("select course_id from enrollments where user_id = $1", userId).Scan(&courses)
	if err != nil {
		return "", nil, err
	}
	rows, err := u.db.Query("select id,title,description from  courses where id=$1", &courses)
	if err != nil {
		return "", nil, err
	}
	course := []models.Courses{}
	for rows.Next() {
		cour := models.Courses{}
		err := rows.Scan(&cour.ID, &cour.Title, &cour.Description)
		if err != nil {
			return "", nil, err
		}
		course = append(course, cour)
	}
	return userId, course, nil
}
// 4-TASK
func (u *User) SearchUser(name, email string) ([]models.Users, error) {

	rows, err := u.db.Query("select user_id,name,email from users where name=$1 or email=$2", &name, &email)
	if err != nil {
		return nil, err
	}
	users := []models.Users{}
	for rows.Next() {
		user := models.Users{}
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
