package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	RegisterProductRoutes(r)
	RegisterUserRoutes(r)
}
