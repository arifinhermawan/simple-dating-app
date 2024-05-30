package swipe

import (
	"context"
	"database/sql"
	"time"

	"github.com/arifinhermawan/simple-dating-app/internal/repository/pgsql"
)

type infraProvider interface {
	GetTimeNow() time.Time
	GetMidnight(input time.Time) time.Time
}

type dbProvider interface {
	BeginTX(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)

	CreateSwipeHistoryInDB(ctx context.Context, req pgsql.CreateSwipeHistoryReq) error
	GetTodaysSwipedListFromDB(ctx context.Context, req pgsql.GetTodaysSwipeListReq) ([]int64, error)
	UpdateSwipeCountInDB(ctx context.Context, req pgsql.UpdateSwipeCountReq) error
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
