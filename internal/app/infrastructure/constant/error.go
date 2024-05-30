package constant

import "errors"

var (
	// error related with data
	ErrAccountNotExist        = errors.New("account not exist")
	ErrDuplicateKey           = errors.New("key already exist")
	ErrPremiumPackageNotExist = errors.New("premium package not exist")
	ErrSwipeLimitReached      = errors.New("swipe limit reached")

	// error related with validation
	ErrEmptyDirection      = errors.New("direction is empty")
	ErrEmptyPremiumPackage = errors.New("premium package is empty")
	ErrEmptyUserID         = errors.New("user id is empty")
	ErrInvalidDirection    = errors.New("invalid direction")
	ErrPasswordEmpty       = errors.New("password is empty")
	ErrUsernameEmpty       = errors.New("username is empty")
	ErrWrongPassword       = errors.New("password is incorrect")
)
