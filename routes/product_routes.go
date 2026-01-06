package routes

import (
	"backendAf/controllers"
	"backendAf/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(r *gin.Engine) {
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	p := r.Group("/products")

	// Public routes
	p.GET("", controllers.GetProducts)
	p.GET("/:id/images", controllers.GetProductImages)
	p.GET("/:id", controllers.GetProduct)

	// Admin sub-group
	adminRoutes := p.Group("")
	adminRoutes.Use(middleware.AuthRequired(), middleware.RequireAdmin())
	{
		adminRoutes.POST("", controllers.CreateProduct)
		adminRoutes.PUT("/:id", controllers.UpdateProduct)
		adminRoutes.DELETE("/:id", controllers.DeleteProduct)
		adminRoutes.POST("/:id/images", controllers.UploadProductImage)
		adminRoutes.PUT("/:id/images/:imageId", controllers.UpdateProductImage)
		adminRoutes.DELETE("/:id/images/:imageId", controllers.DeleteProductImage)
	}
}
