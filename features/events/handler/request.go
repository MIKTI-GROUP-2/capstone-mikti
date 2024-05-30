package handler

type EventInput struct {
	CategoryFK           int    `json:"category_id" form:"category_id" validate:"required"`
	Title                string `json:"event_title" form:"event_title" validate:"required"`
	StartDate            string `json:"start_date" form:"start_date" validate:"required"`
	EndDate              string `json:"end_date" form:"start_date" validate:"required"`
	City                 string `json:"city"  form:"city" validate:"required"`
	StartingPrice        int    `json:"starting_price"  form:"starting_price" validate:"required"`
	Description          string `json:"description"  form:"description" validate:"required"`
	Highlight            string `json:"highlight" form:"highlight"`
	ImportantInformation string `json:"important_information" form:"important_information"`
	Address              string `json:"address"  form:"address" validate:"required"`
	Image                string `json:"image_url" form:"image_url" validate:"required"`
}
