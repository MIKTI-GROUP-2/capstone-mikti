package handler

type EventInput struct {
	CategoryID           int    `json:"category_id" form:"category_id" validate:"required"`
	EventTitle           string `json:"event_title" form:"event_title" validate:"required"`
	StartDate            string `json:"start_date" form:"start_date" validate:"required"`
	EndDate              string `json:"end_date" form:"end_date" validate:"required"`
	City                 string `json:"city"  form:"city" validate:"required"`
	StartingPrice        int    `json:"starting_price"  form:"starting_price" validate:"required"`
	Description          string `json:"description"  form:"description" validate:"required"`
	Highlight            string `json:"highlight" form:"highlight"`
	ImportantInformation string `json:"important_information" form:"important_information"`
	Address              string `json:"address"  form:"address" validate:"required"`
}

type UpdateEvent struct {
	CategoryID           int    `json:"category_id" form:"category_id"`
	EventTitle           string `json:"event_title" form:"event_title"`
	StartDate            string `json:"start_date" form:"start_date"`
	EndDate              string `json:"end_date" form:"end_date"`
	City                 string `json:"city"  form:"city"`
	StartingPrice        int    `json:"starting_price"  form:"starting_price"`
	Description          string `json:"description"  form:"description"`
	Highlight            string `json:"highlight" form:"highlight"`
	ImportantInformation string `json:"important_information" form:"important_information"`
	Address              string `json:"address"  form:"address"`
	ImageUrl             string `json:"image_url"`
	PublicID             string `json:"public_id"`
}
