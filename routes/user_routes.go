package routes

import (
	"backendAf/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine) {
	u := r.Group("/users")
	{
		u.POST("/signup", controllers.SignUp)
		u.POST("/login", controllers.Login)

		// example protected route:
		{
			// r.Use(middleware.AuthRequired())
			r.GET("/me", controllers.Me) // optional: returns current user
		}
	}
}
