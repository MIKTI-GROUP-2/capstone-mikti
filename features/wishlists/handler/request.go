package handler

// Request for Create
type WishlistRequest struct {
	UserID  int `json:"user_id"`
	EventID int `json:"event_id"`
}
