package server

import (
	"github.com/arifinhermawan/simple-dating-app/internal/usecase/account"
	"github.com/arifinhermawan/simple-dating-app/internal/usecase/premium"
	"github.com/arifinhermawan/simple-dating-app/internal/usecase/swipe"
)

type UseCase struct {
	Account *account.UseCase
	Premium *premium.UseCase
	Swipe   *swipe.UseCase
}

func NewUseCase(infra *Infra, svc *Service) *UseCase {
	return &UseCase{
		Account: account.NewUsecase(svc.Account),
		Premium: premium.NewUseCase(svc.Account, svc.Premium, infra),
		Swipe:   swipe.NewUseCase(svc.Account, svc.Swipe),
	}
}
