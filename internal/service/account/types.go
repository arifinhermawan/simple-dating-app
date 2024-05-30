package account

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type UserAccount struct {
	ID       int64
	Username string
	Password string
}

type CreateUserAccountReq struct {
	Username string
	Password string
}

type Token struct {
	Value     string
	ExpiresAt time.Time
}

type claims struct {
	Username string
	jwt.StandardClaims
}
