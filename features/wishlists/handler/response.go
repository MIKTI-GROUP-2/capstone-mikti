package handler

// Response for Create
type CreateWishlistResponse struct {
	UserID  int `json:"user_id"`
	EventID int `json:"event_id"`
}

type CategoryResponse struct {
	ID           uint   `json:"id"`
	CategoryName string `json:"name"`
}

type EventResponse struct {
	ID                   uint             `json:"id"`
	Category             CategoryResponse `json:"category"`
	EventTitle           string           `json:"event_title"`
	StartDate            string           `json:"start_date"`
	EndDate              string           `json:"end_date"`
	City                 string           `json:"city"`
	StartingPrice        int              `json:"starting_price"`
	Description          string           `json:"description"`
	Highlight            string           `json:"highlight"`
	ImportantInformation string           `json:"important_information"`
	Address              string           `json:"address"`
	ImageUrl             string           `json:"image_url"`
}

type WishlistResponse struct {
	ID    uint          `json:"id"`
	Event EventResponse `json:"event"`
}
