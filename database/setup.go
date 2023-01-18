package database

import (
	"fmt"

	"github.com/Rickykn/assignment-2-hactiv8.git/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() (err error) {

	HOST := "localhost"
	PORT := "5432"
	DB_USER := "postgres"
	DB_PASS := "postgres"
	DB_NAME := "pos_db"

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable",
		HOST, DB_USER, DB_PASS, DB_NAME, PORT)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&models.Item{})
	db.AutoMigrate(&models.Order{})

	if err != nil {
		panic(err)
	}

	return nil
}

func Get() *gorm.DB {
	return db
}
