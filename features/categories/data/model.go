package data

import (
	"capstone-mikti/features/events/data"

	"gorm.io/gorm"
)

type Category struct {
	*gorm.Model
	CategoryName string       `gorm:"column:category_name;type:varchar(255)"`
	Event        []data.Event `gorm:"foreignKey:CategoryID"`
}
