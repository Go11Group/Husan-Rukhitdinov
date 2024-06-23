package handler

import (
	"database/sql"
	"pro/storage/postgres"
)

type Handler struct {
	db          *sql.DB
	Card        *postgres.CardRepo
	Station     *postgres.StationRepo
	Terminal    *postgres.TerminalRepo
	Transaction *postgres.TransactionRepo
}

func NewHandler(db *sql.DB, station *postgres.StationRepo, card *postgres.CardRepo, terminal *postgres.TerminalRepo, transaction *postgres.TransactionRepo) *Handler {
	return &Handler{
		db:          db,
		Card:        card,
		Station:     station,
		Terminal:    terminal,
		Transaction: transaction,
	}
}
