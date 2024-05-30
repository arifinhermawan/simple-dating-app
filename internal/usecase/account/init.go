package account

import (
	"context"

	"github.com/arifinhermawan/simple-dating-app/internal/service/account"
)

type accountServiceProvider interface {
	CreateUserAccount(ctx context.Context, username string, password string) error
	GenerateToken(username string) (account.Token, error)
	GetUserAccountByUsername(ctx context.Context, username string) (account.UserAccount, error)
	ValidatePassword(hashed string, password string) bool
}

type UseCase struct {
	account accountServiceProvider
}

func NewUsecase(account accountServiceProvider) *UseCase {
	return &UseCase{
		account: account,
	}
}
