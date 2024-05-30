package constant

import "errors"

var (
	ErrDuplicateKey    = errors.New("key already exist")
	ErrAccountNotExist = errors.New("account not exist")
	ErrWrongPassword   = errors.New("password is incorrect")
)
