package account

import (
	"context"
	"log"
	"strings"

	"github.com/arifinhermawan/simple-dating-app/internal/app/infrastructure/constant"
)

func (u *UseCase) CreateUserAccount(ctx context.Context, username string, password string) error {
	err := u.account.CreateUserAccount(ctx, strings.ToLower(username), password)
	if err != nil {
		metadata := map[string]interface{}{
			"username": username,
		}

		log.Printf("[CreateUserAccount] u.account.CreateUserAccount() got error: %+v\nMeta:%+v\n", err, metadata)
		return err
	}

	return nil
}

func (u *UseCase) Login(ctx context.Context, username string, password string) (Token, error) {
	metadata := map[string]interface{}{
		"username": username,
	}

	username = strings.ToLower(username)
	account, err := u.account.GetUserAccountByUsername(ctx, username)
	if err != nil {
		log.Printf("[Login] u.account.GetUserAccountByUsername() got error: %+v\nMeta:%+v\n", err, metadata)
		return Token{}, err
	}

	if account.ID == 0 {
		return Token{}, constant.ErrAccountNotExist
	}

	isPasswordMatch := u.account.ValidatePassword(account.Password, password)
	if !isPasswordMatch {
		return Token{}, constant.ErrWrongPassword
	}

	token, err := u.account.GenerateToken(username)
	if err != nil {
		log.Printf("[Login] u.account.GenerateToken() got error: %+v\nMeta:%+v\n", err, metadata)
		return Token{}, err
	}

	return Token(token), nil
}
