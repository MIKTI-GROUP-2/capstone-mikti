package handler

import "capstone-mikti/helper/customtime"

type EventResponse struct {
	CategoryFK    int                   `json:"category_id"`
	Title         string                `json:"event_title"`
	StartDate     customtime.CustomTime `json:"start_date"`
	EndDate       customtime.CustomTime `json:"end_date"`
	Description   string                `json:"description"`
	StartingPrice int                   `json:"starting_price"`
}
