package account

// -------------------------
// | structs for parameter |
// -------------------------
type createAccountParam struct {
	Username string `json:"username"`
	Password string `json:"password"`
	PhotoURL string `json:"photo_url"`
}

type loginParam struct {
	Username string `json:"username"`
	Password string `json:"password"`
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

type getProfileResponse struct {
	UserID     int64  `json:"user_id"`
	Username   string `json:"username"`
	PhotoURL   string `json:"photo_url"`
	IsVerified bool   `json:"is_verified"`
}
