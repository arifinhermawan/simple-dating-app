package pgsql

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

func (r *Repository) CreateProfileInDB(ctx context.Context, req CreateProfileReq) error {
	cfg := r.infra.GetConfig().Database
	timeout := time.Duration(cfg.DefaultTimeout) * time.Second
	ctxQuery, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	_, err := req.Tx.ExecContext(ctxQuery, queryCreateProfileInDB, req.UserID, req.Username, req.PhotoURL)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetProfileByUserIDFromDB(ctx context.Context, userID int64) (Profile, error) {
	cfg := r.infra.GetConfig().Database
	timeout := time.Duration(cfg.DefaultTimeout) * time.Second
	ctxQuery, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	var result Profile
	err := r.db.QueryRowContext(ctxQuery, queryGetProfileByUserIDFromDB, userID).
		Scan(
			&result.UserID,
			&result.Username,
			&result.PhotoURL,
			&result.IsVerified,
			&result.IsInfiniteScroll,
			&result.SwipeCount,
			&result.LastSwipeAt,
		)
	if err != nil && err != sql.ErrNoRows {
		return Profile{}, err
	}

	return result, nil
}

func (r *Repository) GetSwappableProfileFromDB(ctx context.Context, userIDs []int64) ([]Profile, error) {
	cfg := r.infra.GetConfig().Database
	timeout := time.Duration(cfg.DefaultTimeout) * time.Second
	ctxQuery, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	param := ""
	format := "$%d"
	for i := 1; i <= len(userIDs); i++ {
		param += fmt.Sprintf(format, i)
		if i != len(userIDs) {
			param += ","
		}
	}

	query := fmt.Sprintf(queryGetSwappableProfileFromDB, param)
	variadic := make([]interface{}, 0, len(userIDs))
	for _, val := range userIDs {
		variadic = append(variadic, val)
	}

	rows, err := r.db.QueryContext(ctxQuery, query, variadic...)
	if err != nil {
		return nil, err
	}

	var result []Profile
	for rows.Next() {
		var profile Profile
		err := rows.Scan(
			&profile.UserID,
			&profile.Username,
			&profile.PhotoURL,
			&profile.IsVerified,
		)
		if err != nil {
			return nil, err
		}

		result = append(result, profile)
	}

	return result, nil
}

func (r *Repository) UpdateProfilePremiumPackageInDB(ctx context.Context, req UpdateProfilePremiumPackageReq) error {
	cfg := r.infra.GetConfig().Database
	timeout := time.Duration(cfg.DefaultTimeout) * time.Second
	ctxQuery, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	_, err := r.db.ExecContext(ctxQuery, queryUpdateProfilePremiumPackageInDB, req.IsVerified, req.IsInfiniteScroll, req.UserID)
	if err != nil {
		return err
	}

	return nil
}
