package server

import (
	"github.com/arifinhermawan/simple-dating-app/internal/repository/pgsql"
	"github.com/arifinhermawan/simple-dating-app/internal/service/account"
)

type Service struct {
	Account *account.Service
}

func NewService(db *pgsql.Repository) *Service {
	return &Service{
		Account: account.NewService(db),
	}
}
