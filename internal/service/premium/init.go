package premium

import (
	"context"

	"github.com/arifinhermawan/simple-dating-app/internal/repository/pgsql"
)

type dbProvider interface {
	CreatePurchaseHistoryInDB(ctx context.Context, userID int64, premiumTypes int) error
	UpdateProfilePremiumPackageInDB(ctx context.Context, req pgsql.UpdateProfilePremiumPackageReq) error
}

type Service struct {
	db dbProvider
}

func NewService(db dbProvider) *Service {
	return &Service{
		db: db,
	}
}
