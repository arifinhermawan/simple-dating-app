package utils

import "github.com/dgrijalva/jwt-go"

type claims struct {
	UserID int64
	jwt.StandardClaims
}
