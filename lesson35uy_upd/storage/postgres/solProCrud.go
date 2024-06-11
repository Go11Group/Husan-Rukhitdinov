package postgres

import (
	"database/sql"
	"my_pro/model"
)

type CrudSolvedProblemRepo struct {
	Db *sql.DB
}

func NewSolvedProblemsRepo(db *sql.DB) *CrudSolvedProblemRepo {
	return &CrudSolvedProblemRepo{Db: db}
}

func (c *CrudSolvedProblemRepo) CreateSolvedProblems(solvedProb model.SolvedProblem) error {
	_, err := c.Db.Exec("insert into solved_problems(user_id, problem_id,solved_at) values($1,$2,$3)",
		&solvedProb.UserID, &solvedProb.ProblemID, &solvedProb.SolvedAt)
	if err != nil {
		return err
	}
	return nil
}

func (u *CrudSolvedProblemRepo) UpdateSolvedProblems(solvedProb model.SolvedProblem, id string) error {
	_, err := u.Db.Exec("update solved_problems set user_id=$1, problem_id=$2, solved_at=$3, where id=$4",
		&solvedProb.UserID, &solvedProb.ProblemID, &solvedProb.SolvedAt, &id)
	if err != nil {
		return err
	}
	return nil
}

func (d *CrudSolvedProblemRepo) DeleteSolvedProblems(id string) error {
	_, err := d.Db.Exec("delete from solved_problems where id=$1", &id)
	if err != nil {
		return err
	}
	return nil
}

func (re *CrudSolvedProblemRepo) ReadSolvedProblems() (error, []model.SolvedProblem) {
	row, err := re.Db.Query("select * from solved_problems")
	if err != nil {
		return err, nil
	}
	defer row.Close()

	var sproblems []model.SolvedProblem
	for row.Next() {
		var sol_problem model.SolvedProblem
		err = row.Scan(&sol_problem.ID, &sol_problem.UserID, &sol_problem.ProblemID)
		if err != nil {
			return err, nil
		}

		sproblems = append(sproblems, sol_problem)
	}
	return nil, sproblems
}
