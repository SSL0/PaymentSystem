package service

import (
	"PaymentAPI/pkg/entity"
	"PaymentAPI/pkg/repository"
	"errors"
)

type PaymentService struct {
	repo *repository.Repository
}

func NewPaymentService(repo *repository.Repository) *PaymentService {
	return &PaymentService{repo}
}

func (s *PaymentService) GetBalance(address string) (float32, error) {
	wallet, err := s.repo.GetWalletByAddress(address)

	if err != nil {
		return -1.0, err
	}

	return wallet.Balance, err
}

func (s *PaymentService) Send(from string, to string, amount float32) error {
	if from == "" || to == "" {
		return errors.New("both addresses are required to complete the transaction")
	}

	if amount <= 0 {
		return errors.New("amount must be positive")
	}

	err := s.repo.TransferMoney(from, to, amount)

	if err != nil {
		return err
	}

	err = s.repo.CreateTransaction(from, to, amount)
	return err
}

func (s *PaymentService) GetLast(count int) ([]entity.Transaction, error) {
	if count <= 0 {
		return []entity.Transaction{}, errors.New("count must be positive")
	}
	transactions, err := s.repo.GetLastNTransactions(count)

	if err != nil {
		return []entity.Transaction{}, err
	}

	return transactions, nil
}
