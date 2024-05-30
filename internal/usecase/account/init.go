package account

import (
	"context"

	"github.com/arifinhermawan/simple-dating-app/internal/service/account"
)

type accountServiceProvider interface {
	CreateUserAccount(ctx context.Context, req account.CreateUserAccountReq) error
	GenerateToken(userID int64) (account.Token, error)
	GetProfileByUserID(ctx context.Context, userID int64) (account.Profile, error)
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
