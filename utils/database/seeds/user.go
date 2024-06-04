package seeds

import (
	"capstone-mikti/features/users"
	"capstone-mikti/helper/enkrip"

	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, username, email, phone_number string) error {
	enkrip := enkrip.New()
	hashPass, _ := enkrip.HashPassword("password")
	return db.Create(&users.User{
		Username:    username,
		Email:       email,
		PhoneNumber: phone_number,
		Password:    hashPass,
		IsAdmin:     true,
		Status:      true,
	}).Error
}
