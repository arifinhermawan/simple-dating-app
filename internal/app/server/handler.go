package server

import (
	"github.com/arifinhermawan/simple-dating-app/internal/handler/account"
	"github.com/arifinhermawan/simple-dating-app/internal/handler/premium"
)

type Handler struct {
	Account *account.Handler
	Premium *premium.Handler
}

func NewHandler(infra *Infra, uc *UseCase) *Handler {
	return &Handler{
		Account: account.NewHandler(uc.Account, infra),
		Premium: premium.NewHandler(infra, uc.Premium),
	}
}
