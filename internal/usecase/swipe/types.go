package swipe

// ------------------
// | request struct |
// ------------------
type SwipeReq struct {
	UserID    int64
	SwipedID  int64
	Direction string
}

// -------------------
// | response struct |
// -------------------
type Profile struct {
	UserID     int64
	Username   string
	PhotoURL   string
	IsVerified bool
}
