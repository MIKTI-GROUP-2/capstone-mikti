package handler

type EventResponse struct {
	CategoryFK    int    `json:"category_id"`
	Title         string `json:"event_title"`
	StartDate     string `json:"start_date"`
	EndDate       string `json:"end_date"`
	Description   string `json:"description"`
	StartingPrice int    `json:"starting_price"`
}
