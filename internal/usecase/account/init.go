package account

import "context"

type accountServiceProvider interface {
	CreateUserAccount(ctx context.Context, username string, password string) error
}

type UseCase struct {
	account accountServiceProvider
}

func NewUsecase(account accountServiceProvider) *UseCase {
	return &UseCase{
		account: account,
	}
}
