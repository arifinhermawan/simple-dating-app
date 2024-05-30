package swipe

import (
	"context"
	"log"

	"github.com/arifinhermawan/simple-dating-app/internal/repository/pgsql"
)

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

		log.Printf("[GetTodaysSwipeList] svc.db.GetTodaysSwipeListFromDB() got an error: %+v\nMeta: %+v\n", err, metadata)
		return nil, err
	}

	return result, nil
}
