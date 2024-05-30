package swipe

import (
	"context"
	"log"

	"github.com/arifinhermawan/simple-dating-app/internal/repository/pgsql"
)

func (svc *Service) CreateSwipeHistory(ctx context.Context, req CreateSwipeHistoryReq) error {
	metadata := map[string]interface{}{
		"user_id":   req.UserID,
		"swiped_id": req.SwipedID,
		"direction": req.Direction,
	}

	tx, err := svc.db.BeginTX(ctx, nil)
	if err != nil {

		log.Printf("[CreateSwipeHistory] svc.db.BeginTx() got an error: %+v\nMeta: %+v\n", err, metadata)
		return err
	}

	defer func() {
		if err != nil {
			if err := tx.Rollback(); err != nil {
				log.Printf("[CreateSwipeHistory] tx.Rollback() got an error: %+v\nMeta: %+v\n", err, metadata)
			}
		}
	}()

	err = svc.db.CreateSwipeHistoryInDB(ctx, pgsql.CreateSwipeHistoryReq{
		Tx:        tx,
		UserID:    req.UserID,
		SwipedID:  req.SwipedID,
		Direction: req.Direction,
	})
	if err != nil {
		log.Printf("[CreateSwipeHistory] svc.db.CreateSwipeHistoryInDB() got an error: %+v\nMeta: %+v\n", err, metadata)
		return err
	}

	err = svc.db.UpdateSwipeCountInDB(ctx, pgsql.UpdateSwipeCountReq{
		Tx:          tx,
		Count:       req.SwipeCount,
		UserID:      req.UserID,
		LastSwipeAt: req.SwipeTime,
	})
	if err != nil {
		log.Printf("[CreateSwipeHistory] svc.db.UpdateSwipeCountInDB() got an error: %+v\nMeta: %+v\n", err, metadata)
		return err
	}

	err = tx.Commit()
	if err != nil {
		log.Printf("[CreateSwipeHistory] tx.Commit() got an error: %+v\nMeta: %+v\n", err, metadata)
	}

	return nil
}

func (svc *Service) GetTodaysSwipedList(ctx context.Context, userID int64) ([]int64, error) {
	timeNow := svc.infra.GetTimeNow()
	midnight := svc.infra.GetMidnight(timeNow)

	result, err := svc.db.GetTodaysSwipedListFromDB(ctx, pgsql.GetTodaysSwipeListReq{
		UserID:    userID,
		StartTime: midnight,
		EndTime:   timeNow,
	})
	if err != nil {
		metadata := map[string]interface{}{
			"user_id": userID,
		}

		log.Printf("[GetTodaysSwipedList] svc.db.GetTodaysSwipedListFromDB() got an error: %+v\nMeta: %+v\n", err, metadata)
		return nil, err
	}

	return result, nil
}
