package pgsql

import "time"

// ------------------
// | request struct |
// ------------------
type GetTodaysSwipeListReq struct {
	UserID    int64
	StartTime time.Time
	EndTime   time.Time
}
