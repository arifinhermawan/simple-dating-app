package server

import "github.com/arifinhermawan/simple-dating-app/internal/handler/account"

type Handler struct {
	Account *account.Handler
}

func NewHandler(uc *UseCase, infra *Infra) *Handler {
	return &Handler{
		Account: account.NewHandler(uc.Account, infra),
	}
}
