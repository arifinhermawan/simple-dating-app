package constant

import "errors"

var (
	// error related with data existence
	ErrAccountNotExist        = errors.New("account not exist")
	ErrDuplicateKey           = errors.New("key already exist")
	ErrPremiumPackageNotExist = errors.New("premium package not exist")

	// error related with validation
	ErrEmptyPremiumPackage = errors.New("premium package is empty")
	ErrPasswordEmpty       = errors.New("password is empty")
	ErrUsernameEmpty       = errors.New("username is empty")
	ErrWrongPassword       = errors.New("password is incorrect")
)
