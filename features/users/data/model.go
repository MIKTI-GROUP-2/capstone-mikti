package data

import (
	"capstone-mikti/features/payments/data"
	"time"

	"gorm.io/gorm"
)

type User struct {
	*gorm.Model
	Username    string         `gorm:"column:username;type:varchar(255)"`
	Email       string         `gorm:"column:email;type:varchar(255)"`
	PhoneNumber string         `gorm:"column:phone_number;type:varchar(255)"`
	Password    string         `gorm:"column:password;type:varchar(255)"`
	IsAdmin     bool           `gorm:"column:is_admin;type:bool"`
	Status      bool           `gorm:"column:status;type:bool"`
	Payment     []data.Payment `gorm:"foreignKey:UserID`
}

type UserResetPass struct {
	*gorm.Model
	Username  string    `gorm:"column:username;type:varchar(255)"`
	Code      string    `gorm:"column:code;type:varchar(255)"`
	ExpiredAt time.Time `gorm:"column:expired_at;type:timestamp"`
}

type UserVerification struct {
	Username  string    `gorm:"column:username;type:varchar(255)"`
	Code      string    `gorm:"column:code;type:varchar(255)"`
	ExpiredAt time.Time `gorm:"column:expired_at;type:timestamp"`
}
