package account

import (
	"context"
	"log"
	"time"

	"github.com/arifinhermawan/simple-dating-app/internal/repository/pgsql"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var (
	generateFromPassword = bcrypt.GenerateFromPassword
	compareHashPassword  = bcrypt.CompareHashAndPassword
)

func (svc *Service) CreateUserAccount(ctx context.Context, username string, password string) error {
	metadata := map[string]interface{}{
		"username": username,
	}

	tx, err := svc.db.BeginTX(ctx, nil)
	if err != nil {
		log.Printf("[CreateUserAccount] svc.db.BeginTX() got an error: %+v\nMeta: %+v\n", err, metadata)
		return err
	}

	defer func() {
		if err != nil {
			if err := tx.Rollback(); err != nil {
				log.Printf("[CreateUserAccount] tx.Rollback() got error")
			}
		}
	}()

	hashed, err := generateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("[CreateUserAccount] generateFromPassword() got error: %+v\nMeta:%+v\n", err, metadata)
		return err
	}

	returningID, err := svc.db.CreateUserAccountInDB(ctx, pgsql.CreateUserReq{
		Tx:       tx,
		Username: username,
		Password: string(hashed),
	})
	if err != nil {
		log.Printf("[CreateUserAccount] svc.db.CreateUserAccountInDB() got error: %+v\nMeta:%+v\n", err, metadata)
		return err
	}

	err = svc.db.CreateProfileInDB(ctx, pgsql.CreateProfileReq{
		Tx:       tx,
		UserID:   returningID,
		Username: username,
	})
	if err != nil {
		log.Printf("[CreateUserAccount] svc.db.CreateProfileInDB() got error: %+v\nMeta:%+v\n", err, metadata)
		return err
	}

	errCommit := tx.Commit()
	if errCommit != nil {
		metadata["user_id"] = returningID
		log.Printf("[CreateUserAccount] tx.Commit() got error: %+v\nMeta:%+v\n", err, metadata)
	}

	return nil
}

func (svc *Service) GetUserAccountByUsername(ctx context.Context, username string) (UserAccount, error) {
	account, err := svc.db.GetUserAccountByUsernameFromDB(ctx, username)
	if err != nil {
		metadata := map[string]interface{}{
			"username": username,
		}
		log.Printf("[GetUserAccountByUsernameFromDB] svc.db.GetUserAccountByUsernameFromDB() got error: %+v\nMeta:%+v\n", err, metadata)
		return UserAccount{}, err
	}

	return UserAccount(account), nil
}

func (svc *Service) GenerateToken(username string) (Token, error) {
	cfg := svc.infra.GetConfig().Token
	expiredAt := time.Now().Add(time.Duration(cfg.DefaultExpiration) * time.Hour)
	claims := claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredAt.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(cfg.Key))
	if err != nil {
		log.Printf("[GenerateToken] jwt.SignedString() got error: %+v\n", err)
		return Token{}, err
	}

	result := Token{
		Value:     tokenString,
		ExpiresAt: expiredAt,
	}

	return result, nil
}

func (svc *Service) ValidatePassword(hashed string, password string) bool {
	err := compareHashPassword([]byte(hashed), []byte(password))
	return err == nil
}
