package account

import (
	"context"
	"database/sql"

	"github.com/arifinhermawan/simple-dating-app/internal/app/infrastructure/configuration"
	"github.com/arifinhermawan/simple-dating-app/internal/repository/pgsql"
)

type dbProvider interface {
	// transaction related method
	BeginTX(ctx context.Context, options *sql.TxOptions) (*sql.Tx, error)

	// user related method
	CreateUserAccountInDB(ctx context.Context, req pgsql.CreateUserReq) (int64, error)
	GetUserAccountByUsernameFromDB(ctx context.Context, username string) (pgsql.UserAccount, error)

	// profile related method
	CreateProfileInDB(ctx context.Context, req pgsql.CreateProfileReq) error
	GetProfileByUserIDFromDB(ctx context.Context, userID int64) (pgsql.Profile, error)
}

type infraProvider interface {
	GetConfig() *configuration.AppConfig
}

type Service struct {
	db    dbProvider
	infra infraProvider
}

func NewService(db dbProvider, infra infraProvider) *Service {
	return &Service{
		db:    db,
		infra: infra,
	}
}
