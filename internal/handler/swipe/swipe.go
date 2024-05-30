package swipe

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/arifinhermawan/simple-dating-app/internal/app/infrastructure/constant"
	"github.com/arifinhermawan/simple-dating-app/internal/usecase/swipe"
)

func (h *Handler) HandlerGetSwappableProfileList(w http.ResponseWriter, r *http.Request) {
	result := response{
		Code:    http.StatusInternalServerError,
		Message: "failed!",
	}
	w.Header().Set("Content-Type", "application/json")

	userIDstr := r.Header.Get("user_id")
	userID, err := strconv.ParseInt(userIDstr, 10, 64)
	if err != nil {
		result.Error = err.Error()
		json.NewEncoder(w).Encode(result)
		return
	}

	profiles, err := h.swipe.GetSwappableProfileList(context.Background(), userID)
	if err != nil {
		result.Code = http.StatusInternalServerError
		result.Error = err.Error()
		json.NewEncoder(w).Encode(result)
		return
	}

	profileList := make([]getSwappableProfileListResponse, 0, len(profiles))
	for _, profile := range profiles {
		profileList = append(profileList, getSwappableProfileListResponse(profile))
	}

	result.Message = "success!"
	result.Code = http.StatusOK
	result.Data = profileList

	json.NewEncoder(w).Encode(result)
}

func (h *Handler) HandlerSwipe(w http.ResponseWriter, r *http.Request) {
	result := response{
		Code:    http.StatusBadRequest,
		Message: "failed!",
	}
	w.Header().Set("Content-Type", "application/json")

	bytes, err := h.infra.ReadAll(r.Body)
	if err != nil {
		result.Error = err.Error()
		json.NewEncoder(w).Encode(result)
		return
	}

	var request swipeReq
	err = h.infra.JsonUnmarshal(bytes, &request)
	if err != nil {
		result.Error = err.Error()
		json.NewEncoder(w).Encode(result)
		return
	}

	if request.Direction == "" {
		result.Error = constant.ErrEmptyDirection.Error()
		json.NewEncoder(w).Encode(result)
		return
	}

	if request.Direction != "right" && request.Direction != "left" {
		result.Error = constant.ErrInvalidDirection.Error()
		json.NewEncoder(w).Encode(result)
		return
	}

	if request.SwipedID == 0 {
		result.Error = constant.ErrEmptyUserID.Error()
		json.NewEncoder(w).Encode(result)
		return
	}

	userIDstr := r.Header.Get("user_id")
	userID, err := strconv.ParseInt(userIDstr, 10, 64)
	if err != nil {
		result.Error = err.Error()
		json.NewEncoder(w).Encode(result)
		return
	}

	err = h.swipe.Swipe(context.Background(), swipe.SwipeReq{
		UserID:    userID,
		SwipedID:  request.SwipedID,
		Direction: request.Direction,
	})
	if err != nil {
		result.Code = http.StatusInternalServerError
		result.Error = err.Error()
		json.NewEncoder(w).Encode(result)
		return
	}

	result.Code = http.StatusOK
	result.Message = "success"
	json.NewEncoder(w).Encode(result)
}
