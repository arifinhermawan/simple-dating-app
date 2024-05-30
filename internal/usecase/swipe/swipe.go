package swipe

import (
	"context"
	"log"

	"github.com/arifinhermawan/simple-dating-app/internal/app/infrastructure/constant"
	"github.com/arifinhermawan/simple-dating-app/internal/service/swipe"
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

func (u *UseCase) Swipe(ctx context.Context, req SwipeReq) error {
	metadata := map[string]interface{}{
		"user_id":   req.UserID,
		"swiped_id": req.SwipedID,
		"direction": req.Direction,
	}

	profile, err := u.account.GetProfileByUserID(ctx, req.UserID)
	if err != nil {
		log.Printf("[Swipe] u.account.GetProfileByUserID() got error:%+v\nMetadata:%+v\n", err, metadata)
		return err
	}

	swipeCount := profile.SwipeCount
	lastSwipe := profile.LastSwipeAt

	timeNow := u.infra.GetTimeNow()
	midnight := u.infra.GetMidnight(timeNow)
	if lastSwipe.Before(midnight) {
		swipeCount = 0
	}

	if swipeCount > 10 && !profile.IsInfiniteScroll {
		return constant.ErrSwipeLimitReached
	}

	swipeCount++

	go func() {
		err := u.swipe.CreateSwipeHistory(ctx, swipe.CreateSwipeHistoryReq{
			UserID:     req.UserID,
			SwipedID:   req.SwipedID,
			Direction:  req.Direction,
			SwipeCount: swipeCount,
			SwipeTime:  timeNow,
		})
		if err != nil {
			log.Printf("[Swipe] u.swipe.CreateSwipeHistory() got error:%+v\nMetadata:%+v\n", err, metadata)
		}
	}()

	return nil
}
