package account

// -------------------------
// | structs for parameter |
// -------------------------
type createAccountParam struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// ------------------------
// | structs for response |
// ------------------------
type defaultResponse struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

type createAccountResponse struct {
	defaultResponse
	Message string `json:"message"`
}
