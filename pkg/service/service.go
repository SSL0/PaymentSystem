package service

import (
	"PaymentAPI/pkg/entity"
	"PaymentAPI/pkg/repository"
)

type Payment interface {
	Send(from string, to string, amount float32) error
	GetLast(count int) ([]entity.Transaction, error)
	GetBalance(address string) (float32, error)
}

type Service struct {
	Payment
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Payment: NewPaymentService(repo),
	}
}
