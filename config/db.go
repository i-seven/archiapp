package config

import (
	"backendAf/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DbConInit() {
	var err error
	dsn := os.Getenv("DBCONSTRING")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	log.Println("DB init was successful")

}
func SyncDB() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Product{})
	DB.AutoMigrate(&models.Ceiling{})
	DB.AutoMigrate(&models.ArchiStyle{})
	DB.AutoMigrate(&models.Garage{})
	DB.AutoMigrate(&models.RoofDetail{})
	DB.AutoMigrate(&models.SpecialFeature{})

	log.Println("db syncing was successful")

}
