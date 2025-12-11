package main

import (
	"backendAf/config"
	"backendAf/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnv()
	config.DbConInit()
	config.SyncDB()
}

func main() {

	r := gin.Default()

	// Serve static images
	r.Static("/images", "./storage/images")

	routes.RegisterRoutes(r)

	r.Run(":8080")

}
