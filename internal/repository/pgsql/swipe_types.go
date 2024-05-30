package pgsql

import (
	"database/sql"
	"time"
)

// ------------------
// | request struct |
// ------------------
type CreateSwipeHistoryReq struct {
	Tx        *sql.Tx
	UserID    int64
	SwipedID  int64
	Direction string
}

type GetTodaysSwipeListReq struct {
	UserID    int64
	StartTime time.Time
	EndTime   time.Time
}
