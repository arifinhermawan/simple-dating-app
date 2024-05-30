package account

import "errors"

var (
	errUsernameEmpty = errors.New("username is empty")
	errPasswordEmpty = errors.New("password is empty")
	errUserExist     = errors.New("user already exist!")
)

const (
	usernameKey = "username"
)
