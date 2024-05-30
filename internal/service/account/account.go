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

func (svc *Service) CreateUserAccount(ctx context.Context, req CreateUserAccountReq) error {
	metadata := map[string]interface{}{
		"username": req.Username,
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

	hashed, err := generateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("[CreateUserAccount] generateFromPassword() got error: %+v\nMeta:%+v\n", err, metadata)
		return err
	}

	returningID, err := svc.db.CreateUserAccountInDB(ctx, pgsql.CreateUserReq{
		Tx:       tx,
		Username: req.Username,
		Password: string(hashed),
	})
	if err != nil {
		log.Printf("[CreateUserAccount] svc.db.CreateUserAccountInDB() got error: %+v\nMeta:%+v\n", err, metadata)
		return err
	}

	err = svc.db.CreateProfileInDB(ctx, pgsql.CreateProfileReq{
		Tx:       tx,
		UserID:   returningID,
		Username: req.Username,
		PhotoURL: req.PhotoURL,
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

func (svc *Service) GetProfileByUserID(ctx context.Context, userID int64) (Profile, error) {
	profile, err := svc.db.GetProfileByUserIDFromDB(ctx, userID)
	if err != nil {
		metadata := map[string]interface{}{
			"user_id": userID,
		}
		log.Printf("[GetProfileByUserID] svc.db.GetProfileByUserIDFromDB() got error: %+v\nMeta:%+v\n", err, metadata)
		return Profile{}, err
	}

	return Profile{
		UserID:           profile.UserID,
		Username:         profile.Username,
		PhotoURL:         profile.PhotoURL.String,
		IsVerified:       profile.IsVerified,
		IsInfiniteScroll: profile.IsInfiniteScroll,
		SwipeCount:       profile.SwipeCount,
		LastSwipeAt:      profile.LastSwipeAt.Time,
	}, nil
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

func (svc *Service) GenerateToken(userID int64) (Token, error) {
	cfg := svc.infra.GetConfig().Token
	expiredAt := time.Now().Add(time.Duration(cfg.DefaultExpiration) * time.Hour)
	claims := claims{
		UserID: userID,
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

func (svc *Service) UpdateProfilePremiumPackage(ctx context.Context, req UpdateProfilePremiumPackageReq) error {
	err := svc.db.UpdateProfilePremiumPackageInDB(ctx, pgsql.UpdateProfilePremiumPackageReq(req))
	if err != nil {
		metadata := map[string]interface{}{
			"user_id": req.UserID,
		}
		log.Printf("[UpdateProfilePremiumPackage] svc.db.UpdateProfilePremiumPackageInDB() got error: %+v\nMetadata:%+v\n", err, metadata)

		return err
	}

	return nil
}

func (svc *Service) ValidatePassword(hashed string, password string) bool {
	err := compareHashPassword([]byte(hashed), []byte(password))
	return err == nil
}
