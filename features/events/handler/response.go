package handler

type EventResponse struct {
	CategoryID           int    `json:"category_id"`
	EventTitle           string `json:"event_title"`
	StartDate            string `json:"start_date"`
	EndDate              string `json:"end_date"`
	City                 string `json:"city"`
	StartingPrice        int    `json:"starting_price"`
	Description          string `json:"description"`
	Highlight            string `json:"highlight"`
	ImportantInformation string `json:"important_information"`
	Address              string `json:"address"`
	ImageUrl             string `json:"image_url"`
}

type DetailEventResponse struct {
	ID                   uint `json:"id"`
	Category             CategoryResponse
	EventTitle           string `json:"event_title"`
	StartDate            string `json:"start_date"`
	EndDate              string `json:"end_date"`
	City                 string `json:"city"`
	StartingPrice        int    `json:"starting_price"`
	Description          string `json:"description"`
	Highlight            string `json:"highlight"`
	ImportantInformation string `json:"important_information"`
	Address              string `json:"address"`
	ImageUrl             string `json:"image_url"`
}

type CategoryResponse struct {
	ID           uint   `json:"id"`
	CategoryName string `json:"category_name"`
}

type AllEventResponse struct {
	ID            uint `json:"id"`
	Category      CategoryResponse
	EventTitle    string `json:"event_title"`
	StartDate     string `json:"start_date"`
	EndDate       string `json:"end_date"`
	StartingPrice int    `json:"starting_price"`
	City          string `json:"city"`
}
