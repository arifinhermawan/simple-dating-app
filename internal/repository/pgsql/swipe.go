package pgsql

import (
	"context"
	"time"
)

const (
	dateTimeLayout = "2006-01-02 15:04:05.000"
)

func (r *Repository) CreateSwipeHistoryInDB(ctx context.Context, req CreateSwipeHistoryReq) error {
	cfg := r.infra.GetConfig().Database
	timeout := time.Duration(cfg.DefaultTimeout) * time.Second
	ctxQuery, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	_, err := req.Tx.ExecContext(ctxQuery, queryCreateSwipeHistoryInDB, req.UserID, req.SwipedID, req.Direction)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetTodaysSwipedListFromDB(ctx context.Context, req GetTodaysSwipeListReq) ([]int64, error) {
	cfg := r.infra.GetConfig().Database
	timeout := time.Duration(cfg.DefaultTimeout) * time.Second
	ctxQuery, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	rows, err := r.db.QueryContext(ctxQuery, queryGetTodaysSwipedList, req.UserID, req.StartTime.Format(dateTimeLayout), req.EndTime.Format(dateTimeLayout))
	if err != nil {
		return nil, err
	}

	var result []int64
	for rows.Next() {
		var id int64
		err := rows.Scan(&id)
		if err != nil {
			return nil, err
		}

		result = append(result, id)
	}

	return result, nil
}
