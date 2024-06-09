package crudsolvedproblems

import (
	"database/sql"
	"gorilla/model"
)

type CrudSolvedProblemRepo struct {
	db *sql.DB
}

func NewSolvedProblemsRepo(db *sql.DB) *CrudSolvedProblemRepo {
	return &CrudSolvedProblemRepo{db: db}
}

func (c *CrudSolvedProblemRepo) CreateSolvedProblems(solvedProb model.SolvedProblem) {
	_, err := c.db.Exec("insert into solved_problems(id, user_id, problem_id,solved_at) values($1,$2,$3,$4)",
		&solvedProb.ID, &solvedProb.UserID, &solvedProb.ProblemID, &solvedProb.SolvedAt)
	if err != nil {
		panic(err)
	}
}

func (u *CrudSolvedProblemRepo) UpdateSolvedProblems(solvedProb model.SolvedProblem) {
	_, err := u.db.Exec("update solved_problems set user_id=$1, problem_id=$2, solved_at=$3, where id=$4",
		&solvedProb.UserID, &solvedProb.ProblemID, &solvedProb.SolvedAt, &solvedProb.ID)
	if err != nil {
		panic(err)
	}
}

func (d *CrudSolvedProblemRepo) DeleteSolvedProblems(solvedProb model.SolvedProblem) {
	_, err := d.db.Exec("delete from solved_problems where id=$1", &solvedProb.ID)
	if err != nil {
		panic(err)
	}
}

func (re *CrudSolvedProblemRepo) ReadSolvedProblems(solvedProb model.SolvedProblem) {
	row, err := re.db.Query("select * from solved_problems")
	if err != nil {
		return
	}
	defer row.Close()

	var sproblems []model.SolvedProblem
	for row.Next() {
		var sol_problem model.SolvedProblem
		err = row.Scan(&sol_problem.ID, &sol_problem.UserID, &sol_problem.ProblemID)
		if err != nil {
			return
		}

		sproblems = append(sproblems, sol_problem)
	}
}
