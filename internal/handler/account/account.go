package account

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/arifinhermawan/simple-dating-app/internal/app/infrastructure/constant"
	"github.com/arifinhermawan/simple-dating-app/internal/usecase/account"
	"github.com/gorilla/mux"
)

func (h *Handler) HandlerCreateUserAccount(w http.ResponseWriter, r *http.Request) {
	result := response{
		Message: "failed!",
	}
	w.Header().Set("Content-Type", "application/json")

	bytes, err := h.infra.ReadAll(r.Body)
	if err != nil {
		result.Code = http.StatusBadRequest
		result.Error = err.Error()

		json.NewEncoder(w).Encode(result)
		return
	}

	var request createAccountParam
	err = h.infra.JsonUnmarshal(bytes, &request)
	if err != nil {
		result.Code = http.StatusBadRequest
		result.Error = err.Error()

		json.NewEncoder(w).Encode(result)
		return
	}

	if request.Username == "" {
		result.Code = http.StatusBadRequest
		result.Error = constant.ErrUsernameEmpty.Error()

		json.NewEncoder(w).Encode(result)
		return
	}

	if request.Password == "" {
		result.Code = http.StatusBadRequest
		result.Error = constant.ErrPasswordEmpty.Error()

		json.NewEncoder(w).Encode(result)
		return
	}

	err = h.account.CreateUserAccount(context.Background(), account.CreateUserAccountReq{
		Username: strings.ToLower(request.Username),
		Password: request.Password,
		PhotoURL: request.PhotoURL,
	})
	if err != nil {
		errMsg := err.Error()

		if err == constant.ErrDuplicateKey {
			errMsg = "username already taken"
		}

		result.Code = http.StatusInternalServerError
		result.Error = errMsg

		json.NewEncoder(w).Encode(result)
		return
	}

	result.Message = "success!"
	result.Code = http.StatusCreated
	json.NewEncoder(w).Encode(result)
}

func (h *Handler) HandlerGetProfile(w http.ResponseWriter, r *http.Request) {
	result := response{
		Code:    http.StatusInternalServerError,
		Message: "failed!",
	}

	w.Header().Set("Content-Type", "application/json")

	path := mux.Vars(r)
	userIDstr := path["user_id"]
	userID, err := strconv.ParseInt(userIDstr, 10, 64)
	if err != nil {
		result.Code = http.StatusBadRequest
		result.Error = "user_id not int"
		json.NewEncoder(w).Encode(result)
		return
	}

	profile, err := h.account.GetProfileByUserID(context.Background(), userID)
	if err != nil {
		result.Error = err.Error()
		json.NewEncoder(w).Encode(result)
		return
	}

	result.Code = http.StatusOK
	result.Message = "success"

	response := getProfileResponse{
		UserID:     profile.UserID,
		Username:   profile.Username,
		PhotoURL:   profile.PhotoURL,
		IsVerified: profile.IsVerified,
	}

	result.Data = response

	json.NewEncoder(w).Encode(result)
}

func (h *Handler) HandlerLogin(w http.ResponseWriter, r *http.Request) {
	result := response{
		Code:    http.StatusInternalServerError,
		Message: "failed!",
	}

	w.Header().Set("Content-Type", "application/json")

	var request loginParam
	bytes, err := h.infra.ReadAll(r.Body)
	if err != nil {
		result.Code = http.StatusBadRequest
		result.Error = err.Error()

		json.NewEncoder(w).Encode(result)
		return
	}

	err = h.infra.JsonUnmarshal(bytes, &request)
	if err != nil {
		result.Code = http.StatusBadRequest
		result.Error = err.Error()

		json.NewEncoder(w).Encode(result)
		return
	}

	if request.Username == "" {
		result.Code = http.StatusBadRequest
		result.Error = constant.ErrUsernameEmpty.Error()

		json.NewEncoder(w).Encode(result)
		return
	}

	if request.Password == "" {
		result.Code = http.StatusBadRequest
		result.Error = constant.ErrPasswordEmpty.Error()

		json.NewEncoder(w).Encode(result)
		return
	}

	token, err := h.account.Login(context.Background(), request.Username, request.Password)
	if err != nil {
		result.Code = http.StatusInternalServerError
		result.Error = err.Error()
		json.NewEncoder(w).Encode(result)

		return
	}

	result.Code = http.StatusOK
	result.Message = "success"

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   token.Value,
		Expires: token.ExpiresAt,
		Path:    "/",
	})

	json.NewEncoder(w).Encode(result)
}
