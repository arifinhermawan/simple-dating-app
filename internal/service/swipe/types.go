package swipe

import "time"

// ------------------
// | request struct |
// ------------------
type CreateSwipeHistoryReq struct {
	UserID     int64
	SwipedID   int64
	Direction  string
	SwipeCount int
	SwipeTime  time.Time
}
