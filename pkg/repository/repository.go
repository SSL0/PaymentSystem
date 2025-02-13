package repository

import (
	"PaymentAPI/pkg/entity"
	"database/sql"
	"github.com/jmoiron/sqlx"
)

type Wallet interface {
	GetWalletByAddress(address string) (entity.Wallet, error)
	TransferMoney(from string, to string, amount float32) error
}

type Transaction interface {
	GetLastNTransactions(count int) ([]entity.Transaction, error)
	CreateTransaction(from string, to string, amount float32) error
}

type Repository struct {
	Wallet
	Transaction
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Wallet:      NewWalletRepository(db),
		Transaction: NewTransactionRepository(db),
		db:          db,
	}
}

func (r *Repository) BeginDBTransaction() (*sql.Tx, error) {
	return r.db.Begin()
}
