package premium

// -------------------------
// | structs for parameter |
// -------------------------
type buyPremiumPackageReq struct {
	PremiumPackage string `json:"premium_package"`
}

// ------------------------
// | structs for response |
// ------------------------
type response struct {
	Code    int         `json:"code"`
	Error   string      `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
