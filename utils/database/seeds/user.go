package seeds

import (
	"capstone-mikti/features/users"
	"capstone-mikti/helper/enkrip"

	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, username, email, phone_number string) error {
	var countData int64
	db.Table("users").Where("username = ?", "admin").Count(&countData)

	if countData < 0 {
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

	return nil
}
