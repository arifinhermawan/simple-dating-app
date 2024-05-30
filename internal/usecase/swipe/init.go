package swipe

import (
	"context"
	"time"

	"github.com/arifinhermawan/simple-dating-app/internal/service/account"
	"github.com/arifinhermawan/simple-dating-app/internal/service/swipe"
)

type infraProvider interface {
	GetTimeNow() time.Time
	GetMidnight(time.Time) time.Time
}

type accountServiceProvider interface {
	GetProfileByUserID(ctx context.Context, userID int64) (account.Profile, error)
	GetSwappableProfile(ctx context.Context, userIDs []int64) ([]account.Profile, error)
}

type swipeServiceProvider interface {
	CreateSwipeHistory(ctx context.Context, req swipe.CreateSwipeHistoryReq) error
	GetTodaysSwipedList(ctx context.Context, userID int64) ([]int64, error)
}

type UseCase struct {
	account accountServiceProvider
	swipe   swipeServiceProvider
	infra   infraProvider
}

func NewUseCase(infra infraProvider, account accountServiceProvider, swipe swipeServiceProvider) *UseCase {
	return &UseCase{
		account: account,
		swipe:   swipe,
		infra:   infra,
	}
}
