package main

import (
	"backendAf/controllers"
	"backendAf/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.DbConInit()
	initializers.SyncDB()
}

func main() {
	router := gin.Default()

	router.POST("/signup", controllers.SignUp)
	router.POST("/login", controllers.Login)

	router.GET("/products", controllers.GetProducts)
	router.GET("/products/:id", controllers.GetProduct)

	router.Use(controllers.AuthMiddleware(), controllers.AdminOnly())
	{
		api := router.Group("/api")
		{
			api.POST("/products", controllers.CreateProduct)
			api.GET("/products", controllers.GetProducts)
			api.GET("/products/:id", controllers.GetProduct)
			api.PUT("/products/:id", controllers.UpdateProduct)
			api.DELETE("/products/:id", controllers.DeleteProduct)
		}
	}
	router.Run()
}
