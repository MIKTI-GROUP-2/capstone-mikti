package database

import (
	"capstone-mikti/configs"
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(c *configs.ProgrammingConfig) *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable", c.DBHost, c.DBPort, c.DBName, c.DBUser, c.DBPassword)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		logrus.Error("Terjadi kesalahan pada database, error : ", err.Error())
		return nil
	}

	return db
}
