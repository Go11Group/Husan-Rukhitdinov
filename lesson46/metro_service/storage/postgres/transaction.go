package postgres

import (
	"database/sql"
	pb "google.golang.org/grpc/metro_service/genproto/metro_service"

	"github.com/google/uuid"
)

type TransactionRepo struct {
	Db *sql.DB
}

func NewTransactionRepo(db *sql.DB) *TransactionRepo {
	return &TransactionRepo{Db: db}
}

func (t *TransactionRepo) Create(transac *pb.Transaction) error {
	_, err := t.Db.Exec("insert into transaction (id, card_id, amount, terminal_id) values ($1,$2,$3,$4)",
		uuid.NewString(), transac.CardId, transac.Amount, transac.TerminalId)
	if err != nil {
		return err
	}
	return nil
}

func (t *TransactionRepo) GetById(id string) (*pb.Transaction, error) {
	var transac = pb.Transaction{Id: id}

	err := t.Db.QueryRow("select card_id, amount, terminal_id from transaction where id = $1", id).Scan(&transac.CardId, &transac.Amount, &transac.TerminalId)
	if err != nil {
		return nil, err
	}
	return &transac, err
}

func (t *TransactionRepo) Update(transac *pb.Transaction, id string) error {
	_, err := t.Db.Exec("update transaction set id = $1, card_id = $2, amount = $3,terminal_id = $4", transac.Id, transac.CardId, transac.Amount, transac.TerminalId)
	if err != nil {
		return err
	}
	return nil
}

func (t *TransactionRepo) Delete(id string) error {
	_, err := t.Db.Exec("delete from transaction where id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (s *TransactionRepo) GetAll(f *pb.Transaction) ([]pb.Transaction, error) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
	)
	query := `select id, card_id, amount, terminal_id from transaction `
	filter := ` where true `

	if len(f.CardId) > 0 {
		params["card_id"] = f.CardId
		filter += "and card_id = :card_id "
	}
	if len(f.Amount) > 0 {
		params["amount"] = f.Amount
		filter += "and amount = :amount "
	}
	if len(f.TerminalId) > 0 {
		params["terminal_id"] = f.TerminalId
		filter += "and terminal_id = :terminal_id "
	}

	query = query + filter

	query, arr = ReplaceQueryParam(query, params)
	rows, err := s.Db.Query(query, arr...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transacs []pb.Transaction
	for rows.Next() {
		var transac pb.Transaction
		err = rows.Scan(&transac.CardId, &transac.Amount, transac.TerminalId)
		if err != nil {
			return nil, err
		}
		transacs = append(transacs, transac)
	}
	return transacs, nil
}
