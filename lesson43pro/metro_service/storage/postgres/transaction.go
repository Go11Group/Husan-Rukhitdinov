package postgres

import (
	"database/sql"
	"pro/models"

	"github.com/google/uuid"
)

type TransactionRepo struct {
	Db *sql.DB
}

func NewTransactionRepo(db *sql.DB) *TransactionRepo {
	return &TransactionRepo{Db: db}
}

func (t *TransactionRepo) Create(transac *models.Transaction) error {
	_,err := t.Db.Exec("insert into transaction (id, card_id, amount, terminal_id) values ($1,$2,$3,$4)",
		uuid.NewString(), transac.CardID,transac.Amount,transac.TerminalId)	
		if err != nil {
			return err 
		}
		return nil 
}

func (t *TransactionRepo) GetById(id string) (*models.Transaction, error) {
	var transac = models.Transaction{Id: id}

	err := t.Db.QueryRow("select card_id, amount, terminal_id from transaction where id = $1", id).Scan(&transac.CardID, &transac.Amount,&transac.TerminalId)
	if err != nil {
		return nil, err
	}
	return &transac, err
}

func (t *TransactionRepo) Update(transac *models.Transaction, id string) error {
	_, err := t.Db.Exec("update transaction set id = $1, card_id = $2, amount = $3,terminal_id = $4", transac.Id, transac.CardID, transac.Amount, transac.TerminalId)
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

func (s *TransactionRepo) GetAll(f models.Transaction) ([]models.Transaction, error) {
	var (
		params = make(map[string]interface{})
		arr []interface{}
	)
	query := `select id, card_id, amount, terminal_id from transaction `
	filter := ` where true `

	
	if len(f.CardID) > 0 {
		params["card_id"] = f.CardID
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
	rows, err := s.Db.Query(query,arr...)
	if err != nil {
		return nil,err 
	}
	defer rows.Close()

	var transacs []models.Transaction
	for rows.Next() {
		var transac models.Transaction
		err = rows.Scan(&transac.CardID,&transac.Amount,transac.TerminalId)
		if err != nil {
			return nil, err
		}
		transacs = append(transacs, transac)
	}
	return transacs, nil 
}