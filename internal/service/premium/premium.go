package premium

import (
	"context"
	"log"
)

func (svc *Service) CreatePurchaseHistory(ctx context.Context, userID int64, premiumType int) error {
	err := svc.db.CreatePurchaseHistoryInDB(ctx, userID, premiumType)
	if err != nil {
		metadata := map[string]interface{}{
			"user_id":      userID,
			"premium_type": premiumType,
		}

		log.Printf("[CreatePurchaseHistory] svc.db.CreatePurchaseHistoryInDB() got an error: %+v\nMeta: %+v\n", err, metadata)
		return err
	}

	return err
}
