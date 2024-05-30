package swipe

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
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
