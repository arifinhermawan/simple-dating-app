package swipe

import (
	"context"

	"github.com/arifinhermawan/simple-dating-app/internal/usecase/swipe"
)

type swipeUseCaseProvider interface {
	GetSwappableProfileList(ctx context.Context, userID int64) ([]swipe.Profile, error)
}

type Handler struct {
	swipe swipeUseCaseProvider
}

func NewHandler(swipe swipeUseCaseProvider) *Handler {
	return &Handler{
		swipe: swipe,
	}
}
