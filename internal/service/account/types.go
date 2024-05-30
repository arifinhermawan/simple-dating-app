package account

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// ------------------
// | request struct |
// ------------------
type CreateUserAccountReq struct {
	Username string
	Password string
	PhotoURL string
}

// -------------------
// | response struct |
// -------------------
type UserAccount struct {
	ID       int64
	Username string
	Password string
}

type Profile struct {
	UserID           int64
	Username         string
	PhotoURL         string
	IsVerified       bool
	IsInfiniteScroll bool
	SwipeCount       int
	LastSwipeAt      time.Time
}

type Token struct {
	Value     string
	ExpiresAt time.Time
}

type claims struct {
	UserID int64
	jwt.StandardClaims
}
