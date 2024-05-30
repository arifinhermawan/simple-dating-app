package account

import "time"

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
type Profile struct {
	UserID     int64
	Username   string
	PhotoURL   string
	IsVerified bool
}

type Token struct {
	Value     string
	ExpiresAt time.Time
}
