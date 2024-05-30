package server

import (
	"github.com/arifinhermawan/simple-dating-app/internal/usecase/account"
)

type UseCase struct {
	Account *account.UseCase
}

func NewUseCase(svc *Service) *UseCase {
	return &UseCase{
		Account: account.NewUsecase(svc.Account),
	}
}
