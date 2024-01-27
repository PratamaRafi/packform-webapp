package config

import (
	"fmt"
	"packform-webapp/models"
	"packform-webapp/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDataBase() (*gorm.DB, error) {
	username := utils.Getenv("DATABASE_USERNAME", "postgres")
	password := utils.Getenv("DATABASE_PASSWORD", "1234")
	host := utils.Getenv("DATABASE_HOST", "127.0.0.1")
	port := utils.Getenv("DATABASE_PORT", "5432")
	database := utils.Getenv("DATABASE_NAME", "packform")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, username, password, database, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	if err := db.AutoMigrate(&models.Customer{}, &models.Companies{}, &models.Deliveries{}, &models.OrderItem{}, &models.Orders{}); err != nil {
		// for debug migrate problem
		panic(err.Error())
	}

	return db, err
}
