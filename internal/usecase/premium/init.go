package premium

import (
	"context"

	"github.com/arifinhermawan/simple-dating-app/internal/app/infrastructure/configuration"
	"github.com/arifinhermawan/simple-dating-app/internal/service/account"
)

type premiumServiceProvider interface {
	CreatePurchaseHistory(ctx context.Context, userID int64, premiumType int) error
}

type accountServiceProvider interface {
	UpdateProfilePremiumPackage(ctx context.Context, req account.UpdateProfilePremiumPackageReq) error
}

type infraProvider interface {
	GetConfig() *configuration.AppConfig
}

type UseCase struct {
	account accountServiceProvider
	premium premiumServiceProvider
	infra   infraProvider
}

func NewUseCase(account accountServiceProvider, premium premiumServiceProvider, infra infraProvider) *UseCase {
	return &UseCase{
		account: account,
		premium: premium,
		infra:   infra,
	}
}
