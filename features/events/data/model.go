package data

import "gorm.io/gorm"

type Event struct {
	*gorm.Model
	EventTitle string `gorm:"column:event_title;type:varchar(225)"`
}
