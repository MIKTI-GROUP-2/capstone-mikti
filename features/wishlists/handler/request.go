package handler

// Request for Create
type WishlistRequest struct {
	EventID int `json:"event_id" form:"event_id" validate:"required"`
}
