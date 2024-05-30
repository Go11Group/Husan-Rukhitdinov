package postgres

import (
	"database/sql"
	"go_modul/model"
)

type StudentRepo struct {
	Db *sql.DB
}

func NewStudentRepo(db *sql.DB) *StudentRepo {
	return &StudentRepo{Db: db}
}

func (u *StudentRepo) GetAllStudents() ([]model.User, error) {
	rows, err := u.Db.Query(`select s.id, s.first_name, s.second_name, s.age, s.gender, s.Phone_number,  c.name from students s`)
	if err != nil {
		return nil, err
	}

	var users []model.User
	var user model.User
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.First_name, &user.Second_name, &user.Age, &user.Gender, &user.Phone_number)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (u *StudentRepo) GetByIDStudent(id string) (model.User, error) {
	var user model.User

	err := u.Db.QueryRow(`select s.id, s.name, age, gender, c.name from students s
					left join course c on c.id = s.course_id where s.id = $1`, id).
		Scan(&user.ID, &user.First_name, &user.Second_name, &user.Age, &user.Gender, &user.Phone_number)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (s *StudentRepo) CreateStudent(user model.User) error {

	_, err := s.Db.Exec(`insert into stundets(first_name, second_name, age, gender, phone_number)
	values($1, $2, $3, $4, $5)`, &user.First_name, &user.Second_name, &user.Age, &user.Gender, &user.Phone_number)
	if err != nil{
		return err
	}
	return nil
}

func (s *StudentRepo) UpdateStudent(user model.User) error {

	_, err := s.Db.Exec(`
	UPDATE students
	SET first_name = $1
	WHERE id = $2`, &user.First_name, &user.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *StudentRepo) DeleteStudent(id string) error {
	_, err := s.Db.Exec(`
	delete from students
	WHERE id = $1`, id)
	if err != nil {
		return err
	}

	return nil
}