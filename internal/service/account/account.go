package account

import (
	"context"
	"log"
	"strings"

	"github.com/arifinhermawan/simple-dating-app/internal/repository/pgsql"
	"golang.org/x/crypto/bcrypt"
)

var (
	generateFromPassword = bcrypt.GenerateFromPassword
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

	username = strings.ToLower(username)
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
		log.Printf("[CreateUserAccount]  tx.Commit() got error: %+v\nMeta:%+v\n", err, metadata)
	}

	return nil
}
