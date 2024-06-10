package postgres

import (
	"database/sql"
	"fmt"
	"my_pro/model"
)

type CrudProblemsRepo struct {
	Db *sql.DB
}

func NewProblemsRepo(db *sql.DB) *CrudProblemsRepo {
	return &CrudProblemsRepo{Db: db}
}

func (c *CrudProblemsRepo) CreateProblems(problem model.Problem) error {
	_, err := c.Db.Exec("insert into problems(title,difficultly,description) values($1,$2,$3)",
		&problem.Title, &problem.Difficulty, &problem.Description)
	if err != nil {
		return err
	}
	return nil
}

func (u *CrudProblemsRepo) UpdateProblems(id string, problem model.Problem) error {
	_, err := u.Db.Exec("update problems set title=$1, difficultly=$2, description=$3, where problem_id=$4",
		&problem.Title, &problem.Difficulty, &problem.Description, &id)
	if err != nil {
		return err
	}
	return nil
}

func (d *CrudProblemsRepo) DeleteProblems(id string) error {

	_, err := d.Db.Exec("delete from problems where problem_id=$1", &id)
	if err != nil {
		return err
	}
	return nil
}

func (re *CrudProblemsRepo) ReadProblems() ([]model.Problem, error) {
	row, err := re.Db.Query("select * from problems")
	if err != nil {
		fmt.Println("0000000000",err)
		return nil, err
	}
	defer row.Close()

	var problems []model.Problem
	for row.Next() {
		var problem model.Problem
		err = row.Scan(&problem.ProblemID, &problem.Title, &problem.Difficulty, &problem.Description)
		if err != nil {
			return nil, err
		}

		problems = append(problems, problem)
	}
	return problems, nil
}
