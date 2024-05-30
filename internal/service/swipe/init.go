package swipe

import (
	"context"
	"time"

	"github.com/arifinhermawan/simple-dating-app/internal/repository/pgsql"
)

type infraProvider interface {
	GetTimeNow() time.Time
	GetMidnight(input time.Time) time.Time
}

type dbProvider interface {
	GetTodaysSwipedListFromDB(ctx context.Context, req pgsql.GetTodaysSwipeListReq) ([]int64, error)
}

type Service struct {
	infra infraProvider
	db    dbProvider
}

func NewService(infra infraProvider, db dbProvider) *Service {
	return &Service{
		infra: infra,
		db:    db,
	}
}
