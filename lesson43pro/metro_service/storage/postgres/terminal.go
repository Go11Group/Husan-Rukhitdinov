package postgres

import (
	"database/sql"

	"pro/models"

	"github.com/google/uuid"
)

type TerminalRepo struct {
	Db *sql.DB
}

func NewTerminalRepo(db *sql.DB) *TerminalRepo {
	return &TerminalRepo{Db: db}
}

func (t *TerminalRepo) Create(terminal *models.Terminal) error {
	_, err := t.Db.Exec("insert into terminal(id,station_id) values ($1,$2)",
		uuid.NewString(), terminal.StationId)
	if err != nil {
		return err
	}
	return nil
}

func (t *TerminalRepo) GetById(id string) (*models.Terminal, error) {
	var terminal = models.Terminal{Id: id}

	err := t.Db.QueryRow("select station_id from terminal where id = $1", id).Scan(&terminal.StationId)
	if err != nil {
		return nil, err
	}
	return &terminal, nil
}

func (t *TerminalRepo) Update(tetminal *models.Terminal, id string) error {
	_, err := t.Db.Exec("update terminal set station_id = $1", tetminal.StationId)
	if err != nil {
		return err
	}
	return nil
}

func (t *TerminalRepo) Delete(id string) error {
	_, err := t.Db.Exec("delete from terminal where id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (t *TerminalRepo) GetAll(f models.Terminal) ([]models.Terminal, error) {
	var (
		params = make(map[string]interface{})
		arr []interface{}
	)
	query := `select id, name from terminal `
	filter := ` where true `

	
	if len(f.StationId) > 0 {
		params["station_id"] = f.StationId
		filter += "and station_id = :station_id "
	}

	query = query + filter 

	query, arr = ReplaceQueryParam(query, params)
	rows, err := t.Db.Query(query,arr...)
	if err != nil {
		return nil,err 
	}
	defer rows.Close()

	var terminals []models.Terminal
	for rows.Next() {
		var terminal models.Terminal
		err = rows.Scan(&terminal.StationId)
		if err != nil {
			return nil, err
		}
		terminals = append(terminals, terminal)
	}
	return terminals, nil 
}