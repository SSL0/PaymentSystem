package repository

import (
	"PaymentAPI/pkg/entity"
	"errors"
	"github.com/jmoiron/sqlx"
)

type WalletRepository struct {
	db *sqlx.DB
}

func NewWalletRepository(db *sqlx.DB) *WalletRepository {
	return &WalletRepository{db}
}

func (r *WalletRepository) GetWalletByAddress(address string) (entity.Wallet, error) {
	var result entity.Wallet

	if address == "" {
		return result, errors.New("address must not be empty")
	}

	err := r.db.QueryRowx("SELECT * FROM wallets WHERE address = $1", address).StructScan(&result)
	return result, err
}

func (r *WalletRepository) TransferMoney(from string, to string, amount float32) error {
	if from == "" || to == "" {
		return errors.New("both addresses are required to complete the transaction")
	}

	if amount <= 0 {
		return errors.New("amount must be positive")
	}

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	query := "SET TRANSACTION ISOLATION LEVEL SERIALIZABLE"
	if _, err = tx.Exec(query); err != nil {
		return err
	}

	query = "UPDATE wallets SET balance = balance - $1 WHERE address = $2"
	if _, err = tx.Exec(query, amount, from); err != nil {
		return err
	}

	query = "UPDATE wallets SET balance = balance + $1 WHERE address = $2"
	if _, err = tx.Exec(query, amount, to); err != nil {
		return err
	}

	err = tx.Commit()
	return err
}
