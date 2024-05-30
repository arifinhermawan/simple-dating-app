package account

import (
	"context"
	"database/sql"

	"github.com/arifinhermawan/simple-dating-app/internal/repository/pgsql"
)

type dbProvider interface {
	BeginTX(ctx context.Context, options *sql.TxOptions) (*sql.Tx, error)
	CreateProfileInDB(ctx context.Context, req pgsql.CreateProfileReq) error
	CreateUserAccountInDB(ctx context.Context, req pgsql.CreateUserReq) (int64, error)
}

type Service struct {
	db dbProvider
}

func NewService(db dbProvider) *Service {
	return &Service{
		db: db,
	}
}
