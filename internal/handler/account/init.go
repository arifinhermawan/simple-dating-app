package account

import (
	"context"
	"io"

	"github.com/arifinhermawan/simple-dating-app/internal/usecase/account"
)

type accountUseCaseProvider interface {
	CreateUserAccount(ctx context.Context, req account.CreateUserAccountReq) error
	GetProfileByUserID(ctx context.Context, userID int64) (account.Profile, error)
	Login(ctx context.Context, username string, password string) (account.Token, error)
}

type infraProvider interface {
	JsonMarshal(input interface{}) ([]byte, error)
	JsonUnmarshal(input []byte, dest interface{}) error
	ReadAll(input io.Reader) ([]byte, error)
}

type Handler struct {
	account accountUseCaseProvider
	infra   infraProvider
}

func NewHandler(account accountUseCaseProvider, infra infraProvider) *Handler {
	return &Handler{
		account: account,
		infra:   infra,
	}
}
