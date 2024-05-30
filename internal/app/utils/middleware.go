package utils

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
)

var (
	errUnauthorized = errors.New("unauthorized")
)

func authMiddleware(endpointHandler func(w http.ResponseWriter, r *http.Request), jwtKey string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				http.Error(w, "unauthorized", http.StatusUnauthorized)
				return
			}

			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}

		claim, err := checkToken(cookie.Value, jwtKey)
		if err != nil {
			if err == errUnauthorized {
				http.Error(w, "unauthorized", http.StatusUnauthorized)
				return
			}

			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}

		r.Header.Add("user_id", strconv.FormatInt(claim.UserID, 10))
		endpointHandler(w, r)
	})
}

func checkToken(token string, jwtKey string) (*claims, error) {
	claim := &claims{}

	jwtToken, err := jwt.ParseWithClaims(token, claim, func(t *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, errUnauthorized
		}

		return nil, err
	}

	if !jwtToken.Valid {
		return nil, errUnauthorized
	}

	return claim, nil
}
