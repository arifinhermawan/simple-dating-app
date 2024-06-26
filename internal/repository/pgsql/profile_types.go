package pgsql

import (
	"database/sql"
	"time"
)

// ------------------
// | request struct |
// ------------------
type CreateProfileReq struct {
	Tx       *sql.Tx
	UserID   int64
	Username string
	PhotoURL string
}

type UpdateProfilePremiumPackageReq struct {
	UserID           int64
	IsVerified       bool
	IsInfiniteScroll bool
}

type UpdateSwipeCountReq struct {
	Tx          *sql.Tx
	Count       int
	UserID      int64
	LastSwipeAt time.Time
}

// -------------------
// | response struct |
// -------------------

type Profile struct {
	UserID           int64          `db:"user_id"`
	Username         string         `db:"username"`
	PhotoURL         sql.NullString `db:"photo_url"`
	IsVerified       bool           `db:"is_verified"`
	IsInfiniteScroll bool           `db:"is_infinite_scroll"`
	SwipeCount       int            `db:"swipe_count"`
	LastSwipeAt      sql.NullTime   `db:"last_swipe"`
}
