package account

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/arifinhermawan/simple-dating-app/internal/app/infrastructure/constant"
)

func (h *Handler) HandlerCreateUserAccount(w http.ResponseWriter, r *http.Request) {
	result := createAccountResponse{
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
		result.Error = errUsernameEmpty.Error()

		json.NewEncoder(w).Encode(result)
		return
	}

	if request.Password == "" {
		result.Code = http.StatusBadRequest
		result.Error = errPasswordEmpty.Error()

		json.NewEncoder(w).Encode(result)
		return
	}

	err = h.account.CreateUserAccount(context.Background(), request.Username, request.Password)
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
