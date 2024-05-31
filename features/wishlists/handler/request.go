package handler

type WishlistRequest struct {
	UserID  int `json:"user_id"`
	EventID int `json:"event_id"`
}
