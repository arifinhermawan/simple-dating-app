package server

import (
	"github.com/arifinhermawan/simple-dating-app/internal/handler/account"
	"github.com/arifinhermawan/simple-dating-app/internal/handler/premium"
	"github.com/arifinhermawan/simple-dating-app/internal/handler/swipe"
)

type Handler struct {
	Account *account.Handler
	Premium *premium.Handler
	Swipe   *swipe.Handler
}

func NewHandler(infra *Infra, uc *UseCase) *Handler {
	return &Handler{
		Account: account.NewHandler(uc.Account, infra),
		Premium: premium.NewHandler(infra, uc.Premium),
		Swipe:   swipe.NewHandler(uc.Swipe),
	}
}
