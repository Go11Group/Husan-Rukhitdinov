package postgres

import (
	"database/sql"

	pb "google.golang.org/grpc/metro_service/genproto/metro_service"

	"github.com/google/uuid"
)

type CardRepo struct {
	Db *sql.DB
}

func NewCardRepo(db *sql.DB) *CardRepo {
	return &CardRepo{Db: db}
}

func (c *CardRepo) Create(card *pb.Card) error {
	_, err := c.Db.Exec("insert into card (id,number,user_id) values ($1, $2, $3)",
		uuid.NewString(), card.Name, card.UserId)
	if err != nil {
		return err
	}
	return nil
}

func (c *CardRepo) GetById(id string) (*pb.Card, error) {
	var card = pb.Card{Id: id}

	err := c.Db.QueryRow("select number, user_id from card where id = $1", id).Scan(&card.Name)
	if err != nil {
		return nil, err
	}
	return &card, err
}

func (c *CardRepo) Update(card *pb.Card, id string) error {
	_, err := c.Db.Exec("update card set id = $1, number = $2, user_id = $3", card.Id, card.Name, card.UserId)
	if err != nil {
		return err
	}
	return nil
}

func (c *CardRepo) Delete(id string) error {
	_, err := c.Db.Exec("delete from card where id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (c *CardRepo) GetAllCard(f *pb.Card) ([]pb.Card, error) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
	)
	query := `select id, name from card `
	filter := ` where true `

	if len(f.Name) > 0 {
		params["number"] = f.Name
		filter += "and number = :number "
	}
	if len(f.UserId) > 0 {
		params["user_id"] = f.UserId
		filter += "and user_id = :user_id "
	}

	query = query + filter

	query, arr = ReplaceQueryParam(query, params)
	rows, err := c.Db.Query(query, arr...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cards []pb.Card
	for rows.Next() {
		var card pb.Card
		err = rows.Scan(&card.Name, &card.UserId)
		if err != nil {
			return nil, err
		}
		cards = append(cards, card)
	}
	return cards, nil
}
