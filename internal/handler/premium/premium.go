package premium

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/arifinhermawan/simple-dating-app/internal/app/infrastructure/constant"
)

func (h *Handler) HandlerBuyPremiumPackage(w http.ResponseWriter, r *http.Request) {
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

	var request buyPremiumPackageReq
	err = h.infra.JsonUnmarshal(bytes, &request)
	if err != nil {
		result.Error = err.Error()
		json.NewEncoder(w).Encode(result)
		return
	}

	if request.PremiumPackage == "" {
		result.Error = constant.ErrEmptyPremiumPackage.Error()
		json.NewEncoder(w).Encode(result)
		return
	}

	mapPremiumPackage := h.infra.GetConfig().Premium.MapPackageToID
	if _, ok := mapPremiumPackage[request.PremiumPackage]; !ok {
		result.Error = constant.ErrPremiumPackageNotExist.Error()
		json.NewEncoder(w).Encode(result)
		return
	}

	userIDstr := r.Header.Get("user_id")
	userID, err := strconv.ParseInt(userIDstr, 10, 64)
	if err != nil {
		result.Code = http.StatusInternalServerError
		result.Error = err.Error()
		json.NewEncoder(w).Encode(result)
		return
	}

	premiumPackageID := mapPremiumPackage[request.PremiumPackage]
	err = h.premium.BuyPremiumPackage(context.Background(), userID, premiumPackageID)
	if err != nil {
		result.Code = http.StatusInternalServerError
		result.Error = err.Error()
		json.NewEncoder(w).Encode(result)
		return
	}

	result.Code = http.StatusCreated
	result.Message = "success"
	json.NewEncoder(w).Encode(result)
}
