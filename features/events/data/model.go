package data

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	*gorm.Model
	CategoryID           int       `gorm:"column:category_id;type:integer"`
	EventTitle           string    `gorm:"column:event_title;type:varchar(255)"`
	StartDate            time.Time `gorm:"column:start_date;type:date"`
	EndDate              time.Time `gorm:"column:end_date;type:date"`
	City                 string    `gorm:"column:city;type:varchar(255)"`
	StartingPrice        int       `gorm:"column:starting_price;type:integer"`
	Description          string    `gorm:"column:description;type:text"`
	Highlight            string    `gorm:"column:highlight;type:varchar(255);null"`
	ImportantInformation string    `gorm:"column:important_information;type:varchar(255);null"`
	Address              string    `gorm:"column:address;type:varchar(255)"`
	ImageUrl             string    `gorm:"column:image_url;type:varchar(255);null"`
	PublicID             string    `gorm:"column:public_id;type:varchar(255);null"`
}
