package entity

import "time"

type Transaction struct {
	Id         int
	FromWallet string `db:"from_wallet"`
	ToWallet   string `db:"to_wallet"`
	Amount     float32
	Time       time.Time
}
