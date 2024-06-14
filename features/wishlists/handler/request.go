package handler

// Request for Create
type CreateWishlistRequest struct {
	EventID int `json:"event_id" form:"event_id" validate:"required"`
}
