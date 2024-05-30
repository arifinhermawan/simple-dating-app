package pgsql

import (
	"context"
	"database/sql"
	"time"

	"github.com/arifinhermawan/simple-dating-app/internal/app/infrastructure/constant"
	"github.com/lib/pq"
)

func (r *Repository) GetUserAccountByUsernameFromDB(ctx context.Context, username string) (UserAccount, error) {
	cfg := r.infra.GetConfig().Database
	timeout := time.Duration(cfg.DefaultTimeout) * time.Second
	ctxQuery, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	var id int64
	var pass string
	var uname string
	err := r.db.QueryRowContext(ctxQuery, queryGetUserAccountByUsername, username).Scan(&id, &uname, &pass)
	if err != nil && err != sql.ErrNoRows {
		return UserAccount{}, err
	}

	userAccount := UserAccount{
		ID:       id,
		Username: uname,
		Password: pass,
	}

	return userAccount, nil
}

func (r *Repository) CreateUserAccountInDB(ctx context.Context, req CreateUserReq) (int64, error) {
	cfg := r.infra.GetConfig().Database
	timeout := time.Duration(cfg.DefaultTimeout) * time.Second
	ctxQuery, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	var id int64
	err := req.Tx.QueryRowContext(ctxQuery, queryCreateUserAccount, req.Username, req.Password).Scan(&id)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			if pgErr.Code == "23505" {
				return 0, constant.ErrDuplicateKey
			}
		}
		return 0, err
	}

	return id, nil
}
