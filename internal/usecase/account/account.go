package account

import (
	"context"
	"log"
)

func (u *UseCase) CreateUserAccount(ctx context.Context, username string, password string) error {
	err := u.account.CreateUserAccount(ctx, username, password)
	if err != nil {
		metadata := map[string]interface{}{
			"username": username,
		}

		log.Printf("[CreateUserAccount] u.account.CreateUserAccount() got error: %+v\nMeta:%+v\n", err, metadata)
		return err
	}

	return nil
}
