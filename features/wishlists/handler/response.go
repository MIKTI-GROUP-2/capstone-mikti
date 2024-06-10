package handler

// Response for Create
type WishlistResponse struct {
	UserID  int `json:"user_id"`
	EventID int `json:"event_id"`
}
