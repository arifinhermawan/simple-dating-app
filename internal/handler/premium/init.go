package premium

import (
	"context"
	"io"

	"github.com/arifinhermawan/simple-dating-app/internal/app/infrastructure/configuration"
)

type infraProvider interface {
	GetConfig() *configuration.AppConfig
	JsonMarshal(input interface{}) ([]byte, error)
	JsonUnmarshal(input []byte, dest interface{}) error
	ReadAll(input io.Reader) ([]byte, error)
}

type premiumUseCaseProvider interface {
	BuyPremiumPackage(ctx context.Context, userID int64, premiumPackage int) error
}

type Handler struct {
	infra   infraProvider
	premium premiumUseCaseProvider
}

func NewHandler(infra infraProvider, premium premiumUseCaseProvider) *Handler {
	return &Handler{
		infra:   infra,
		premium: premium,
	}
}
