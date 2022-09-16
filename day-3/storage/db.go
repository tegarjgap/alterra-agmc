package storage

import (
	"log"
	"myapp/config"
	"myapp/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitialMigration() {
	DB.AutoMigrate(&models.Users{})
}

func InitDB() *gorm.DB {
	var err error

	DB, err = gorm.Open(mysql.Open(config.GetMySQLConnectionString()), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}

	return DB
}
