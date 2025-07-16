package repository

import (
	"PaymentAPI/pkg/entity"
	"errors"
	"github.com/jmoiron/sqlx"
)

type TransactionRepository struct {
	db *sqlx.DB
}

func NewTransactionRepository(db *sqlx.DB) *TransactionRepository {
	return &TransactionRepository{db}
}

func (r *TransactionRepository) GetLastNTransactions(count int) ([]entity.Transaction, error) {
	var result []entity.Transaction
	if count <= 0 {
		return result, errors.New("count must be positive")
	}

	rows, err := r.db.Queryx("SELECT * FROM transactions ORDER BY time DESC LIMIT $1", count)
	if err != nil {
		return []entity.Transaction{}, err
	}

	for rows.Next() {
		var transaction entity.Transaction

		err := rows.StructScan(&transaction)
		if err != nil {
			return []entity.Transaction{}, err
		}

		result = append(result, transaction)
	}

	return result, err
}

func (r *TransactionRepository) CreateTransaction(from string, to string, amount float32) error {
	if from == "" || to == "" {
		return errors.New("both addresses are required to complete the transaction")
	}

	if amount <= 0 {
		return errors.New("amount must be positive")
	}

	_, err := r.db.Exec("INSERT INTO transactions (from_wallet, to_wallet, amount) VALUES ($1, $2, $3)",
		from, to, amount)
	return err
}
