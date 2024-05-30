package swipe

import (
	"context"
	"log"
)

func (u *UseCase) GetSwappableProfileList(ctx context.Context, userID int64) ([]Profile, error) {
	metadata := map[string]interface{}{
		"user_id": userID,
	}

	swappedIDs, err := u.swipe.GetTodaysSwipedList(ctx, userID)
	if err != nil {
		log.Printf("[GetTodaysSwipeList] u.swipe.GetTodaysSwipeList() got error:%+v\nMetadata:%+v\n", err, metadata)
		return nil, err
	}

	// to exclude own's profile
	swappedIDs = append(swappedIDs, userID)

	profiles, err := u.account.GetSwappableProfile(ctx, swappedIDs)
	if err != nil {
		log.Printf("[GetTodaysSwipeList] u.account.GetSwappableProfile() got error:%+v\nMetadata:%+v\n", err, metadata)
		return nil, err
	}

	result := make([]Profile, 0, len(profiles))
	for _, profile := range profiles {
		profile := Profile{
			UserID:     profile.UserID,
			Username:   profile.Username,
			PhotoURL:   profile.PhotoURL,
			IsVerified: profile.IsVerified,
		}

		result = append(result, profile)
	}

	return result, nil
}
