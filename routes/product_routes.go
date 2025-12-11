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
	{
		p.Use(middleware.RequireAdmin())
		{

			p.POST("/", controllers.CreateProduct)
			p.GET("/", controllers.GetProducts)
			p.GET("/:id", controllers.GetProduct)
			p.PUT("/:id", controllers.UpdateProduct)
			p.DELETE("/:id", controllers.DeleteProduct)

			// image routes (consistent naming: /products/:productId/images/...)
			p.POST("/:id/images", controllers.UploadProductImage)
			p.PUT("/:id/images/:imageId", controllers.UpdateProductImage) // update/replace image
			p.GET("/:id/images", controllers.GetProductImages)
			p.DELETE("/:id/images/:imageId", controllers.DeleteProductImage)
		}
	}
}
