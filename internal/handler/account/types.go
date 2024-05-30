package account

// -------------------------
// | structs for parameter |
// -------------------------
type createAccountParam struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type loginParam struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// ------------------------
// | structs for response |
// ------------------------
type response struct {
	Code    int    `json:"code"`
	Error   string `json:"error"`
	Message string `json:"message"`
}
