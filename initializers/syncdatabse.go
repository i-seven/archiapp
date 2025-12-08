package initializers

import (
	"backendAf/models"
	"log"
)

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
