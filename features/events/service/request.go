package service

import "time"

type EventInput struct {
	ID                   uint      `json:"id"`
	CategoryFK           string    `json:"category_id"`
	Title                string    `json:"event_title"`
	StartDate            time.Time `json:"start_date"`
	EntDate              time.Time `json:"end_date"`
	City                 string    `json:"city"`
	StartingPrice        int       `json:"starting_price"`
	Description          string    `json:"description"`
	Highlight            string    `json:"highlight"`
	ImportantInformation string    `json:"important_information"`
	Address              string    `json:"address"`
	Image                string    `json:"image_url"`
	Password             string    `json:"password" form:"password" validate:"required"`
}
