package premium

import (
	"context"
	"log"

	"github.com/arifinhermawan/simple-dating-app/internal/service/account"
)

func (u *UseCase) BuyPremiumPackage(ctx context.Context, userID int64, premiumPackage int) error {
	metadata := map[string]interface{}{
		"user_id": userID,
	}

	err := u.premium.CreatePurchaseHistory(ctx, userID, premiumPackage)
	if err != nil {
		log.Printf("[BuyPremiumPackage] u.premium.CreatePurchaseHistory() got error: %+v\nMeta: %+v\n", err, metadata)
		return err
	}

	err = u.account.UpdateProfilePremiumPackage(context.Background(), account.UpdateProfilePremiumPackageReq{
		UserID:           userID,
		IsVerified:       premiumPackage == idVerified,
		IsInfiniteScroll: premiumPackage == idInfiniteScroll,
	})
	if err != nil {
		log.Printf("[BuyPremiumPackage] u.account.UpdateProfilePremiumPackage() got error: %+v\nMeta: %+v\n", err, metadata)
		return err
	}

	return nil
}
