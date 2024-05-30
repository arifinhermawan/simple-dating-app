package swipe

import (
	"context"
	"io"

	"github.com/arifinhermawan/simple-dating-app/internal/usecase/swipe"
)

type infraProvider interface {
	JsonMarshal(input interface{}) ([]byte, error)
	JsonUnmarshal(input []byte, dest interface{}) error
	ReadAll(input io.Reader) ([]byte, error)
}

type swipeUseCaseProvider interface {
	GetSwappableProfileList(ctx context.Context, userID int64) ([]swipe.Profile, error)
	Swipe(ctx context.Context, req swipe.SwipeReq) error
}

type Handler struct {
	infra infraProvider
	swipe swipeUseCaseProvider
}

func NewHandler(infra infraProvider, swipe swipeUseCaseProvider) *Handler {
	return &Handler{
		infra: infra,
		swipe: swipe,
	}
}
