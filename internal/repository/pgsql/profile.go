package pgsql

import (
	"context"
	"time"
)

func (r *Repository) CreateProfileInDB(ctx context.Context, req CreateProfileReq) error {
	cfg := r.infra.GetConfig().Database
	timeout := time.Duration(cfg.DefaultTimeout) * time.Second
	ctxQuery, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	_, err := req.Tx.ExecContext(ctxQuery, queryCreateProfileInDB, req.UserID, req.Username)
	if err != nil {
		return err
	}

	return nil
}
