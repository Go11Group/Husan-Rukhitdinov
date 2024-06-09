package crudproblem

import (
	"database/sql"
	"gorilla/model"
)

type CrudProblemsRepo struct {
	db *sql.DB
}

func NewProblemsRepo(db *sql.DB) *CrudProblemsRepo {
	return &CrudProblemsRepo{db: db}
}

func (c *CrudProblemsRepo) CreateProblems(problem model.Problem) {
	_, err := c.db.Exec("insert into problems(problem_id, title,difficultly,description) values($1,$2,$3,$4)",
		&problem.ProblemID, &problem.Title, &problem.Difficulty, &problem.Description)
	if err != nil {
		panic(err)
	}
}

func (u *CrudProblemsRepo) UpdateProblems(problem model.Problem) {
	_, err := u.db.Exec("update problems set title=$1, difficultly=$2, description=$3, where problem_id=$4",
		&problem.Title, &problem.Difficulty, &problem.Description, &problem.ProblemID)
	if err != nil {
		panic(err)
	}
}

func (d *CrudProblemsRepo) DeleteProblems(problem model.Problem) {

	_, err := d.db.Exec("delete from problems where problem_id=$1", &problem.ProblemID)
	if err != nil {
		panic(err)
	}
}

func (re *CrudProblemsRepo) ReadProblems(problem model.Problem) {
	row, err := re.db.Query("select * from problems")
	if err != nil {
		return
	}
	defer row.Close()

	var problems []model.Problem
	for row.Next() {
		var problem model.Problem
		err = row.Scan(&problem.ProblemID, &problem.Title, &problem.Difficulty, &problem.Description)
		if err != nil {
			return
		}

		problems = append(problems, problem)
	}
}
