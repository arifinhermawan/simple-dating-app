package account

import (
	"context"
	"log"
	"strings"

	"github.com/arifinhermawan/simple-dating-app/internal/app/infrastructure/constant"
	"github.com/arifinhermawan/simple-dating-app/internal/service/account"
)

func (u *UseCase) CreateUserAccount(ctx context.Context, req CreateUserAccountReq) error {
	err := u.account.CreateUserAccount(ctx, account.CreateUserAccountReq{
		Username: req.Username,
		Password: req.Password,
		PhotoURL: req.PhotoURL,
	})
	if err != nil {
		metadata := map[string]interface{}{
			"username": req.Username,
		}

		log.Printf("[CreateUserAccount] u.account.CreateUserAccount() got error: %+v\nMeta:%+v\n", err, metadata)
		return err
	}

	return nil
}

func (u *UseCase) GetProfileByUserID(ctx context.Context, userID int64) (Profile, error) {
	metadata := map[string]interface{}{
		"user_id": userID,
	}

	profile, err := u.account.GetProfileByUserID(ctx, userID)
	if err != nil {
		log.Printf("[GetProfileByUserID] u.account.GetProfileByUserID() got error: %+v\nMeta:%+v\n", err, metadata)
		return Profile{}, err
	}

	if profile.UserID == 0 {
		return Profile{}, constant.ErrAccountNotExist
	}

	return Profile{
		UserID:     profile.UserID,
		Username:   profile.Username,
		PhotoURL:   profile.PhotoURL,
		IsVerified: profile.IsVerified,
	}, nil
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

	token, err := u.account.GenerateToken(account.ID)
	if err != nil {
		log.Printf("[Login] u.account.GenerateToken() got error: %+v\nMeta:%+v\n", err, metadata)
		return Token{}, err
	}

	return Token(token), nil
}
