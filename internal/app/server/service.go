package server

import (
	"github.com/arifinhermawan/simple-dating-app/internal/repository/pgsql"
	"github.com/arifinhermawan/simple-dating-app/internal/service/account"
	"github.com/arifinhermawan/simple-dating-app/internal/service/premium"
	"github.com/arifinhermawan/simple-dating-app/internal/service/swipe"
)

type Service struct {
	Account *account.Service
	Premium *premium.Service
	Swipe   *swipe.Service
}

func NewService(db *pgsql.Repository, infra *Infra) *Service {
	return &Service{
		Account: account.NewService(db, infra),
		Premium: premium.NewService(db),
		Swipe:   swipe.NewService(infra, db),
	}
}
