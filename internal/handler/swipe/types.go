package swipe

// -------------------------
// | structs for parameter |
// -------------------------
type swipeReq struct {
	SwipedID  int64  `json:"swiped_id"`
	Direction string `json:"direction"`
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

type getSwappableProfileListResponse struct {
	UserID     int64  `json:"user_id"`
	Username   string `json:"username"`
	PhotoURL   string `json:"photo_url"`
	IsVerified bool   `json:"is_verified"`
}
