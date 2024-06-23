package postgres

import (
	"database/sql"

	"pro/models"

	"github.com/google/uuid"
)

type StationRepo struct {
	Db *sql.DB
}

func NewStationRepo(db *sql.DB) *StationRepo {
	return &StationRepo{Db: db}
}

func (s *StationRepo) Create(station *models.CreateStation) error {

	_, err := s.Db.Exec("insert into station(id, name) values ($1, $2)",
		uuid.NewString(), station.Name)
	if err != nil {

		return err
	}
	return nil
}

func (s *StationRepo) GetById(id string) (*models.Station, error) {
	var station = models.Station{Id: id}

	err := s.Db.QueryRow("select name from station where id = $1", id).
		Scan(&station.Name)
	if err != nil {
		return nil, err
	}

	return &station, nil
}

func (s *StationRepo) Update(station *models.CreateStation, id string) error {
	_, err := s.Db.Exec("update station set name = $1", station.Name)
	if err != nil {
		return err
	}
	return nil
}

func (s *StationRepo) Delete(id string) error {
	_, err := s.Db.Exec("delete from station where id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
func (s *StationRepo) GetAll(f models.CreateStation) ([]models.CreateStation, error) {
	var (
		params = make(map[string]interface{})
		arr []interface{}
	)
	query := `select id, name from station `
	filter := ` where true `

	
	if len(f.Name) > 0 {
		params["name"] = f.Name
		filter += "and name = :name "
	}

	query = query + filter 

	query, arr = ReplaceQueryParam(query, params)
	rows, err := s.Db.Query(query,arr...)
	if err != nil {
		return nil,err 
	}
	defer rows.Close()

	var stations []models.CreateStation
	for rows.Next() {
		var station models.CreateStation
		err = rows.Scan(&station.Name)
		if err != nil {
			return nil, err
		}
		stations = append(stations, station)
	}
	return stations, nil 
}
