package swipe

import (
	"context"

	"github.com/arifinhermawan/simple-dating-app/internal/service/account"
)

type accountServiceProvider interface {
	GetSwappableProfile(ctx context.Context, userIDs []int64) ([]account.Profile, error)
}

type swipeServiceProvider interface {
	GetTodaysSwipedList(ctx context.Context, userID int64) ([]int64, error)
}

type UseCase struct {
	account accountServiceProvider
	swipe   swipeServiceProvider
}

func NewUseCase(account accountServiceProvider, swipe swipeServiceProvider) *UseCase {
	return &UseCase{
		account: account,
		swipe:   swipe,
	}
}
