package data

import "gorm.io/gorm"

type Category struct {
	*gorm.Model
	CategoryName string `gorm:"column:category_name;type:varchar(255)"`
}
