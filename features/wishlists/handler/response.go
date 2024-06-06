package handler

// Response for Create
type WishlistResponse struct {
	ID      uint `json:"id"`
	UserID  int  `json:"user_id"`
	EventID int  `json:"event_id"`
}
