package account

import (
	"context"
	"io"
)

type accountUseCaseProvider interface {
	CreateUserAccount(ctx context.Context, username string, password string) error
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
