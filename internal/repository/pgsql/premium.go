package pgsql

import (
	"context"
	"time"
)

func (r *Repository) CreatePurchaseHistoryInDB(ctx context.Context, userID int64, premiumTypes int) error {
	cfg := r.infra.GetConfig().Database
	timeout := time.Duration(cfg.DefaultTimeout) * time.Second
	ctxQuery, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	_, err := r.db.ExecContext(ctxQuery, queryCreatePurchaseHistory, userID, premiumTypes)
	if err != nil {
		return err
	}

	return nil
}
