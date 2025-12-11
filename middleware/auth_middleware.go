package middleware

import (
	"backendAf/config"
	"backendAf/models"
	"backendAf/repositories"
	"backendAf/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// AuthRequired parses Authorization header "Bearer <token>" and sets user in context.
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		h := c.GetHeader("Authorization")
		if h == "" {
			utils.JSONError(c, http.StatusUnauthorized, "UNAUTHORIZED", "authorization header required")
			c.Abort()
			return
		}
		parts := strings.Fields(h)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			utils.JSONError(c, http.StatusUnauthorized, "UNAUTHORIZED", "invalid authorization format")
			c.Abort()
			return
		}
		tokStr := parts[1]

		token, err := jwt.Parse(tokStr, func(token *jwt.Token) (interface{}, error) {
			// validate signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrTokenMalformed
			}
			return config.JWTSecret, nil
		})
		if err != nil || !token.Valid {
			utils.JSONError(c, http.StatusUnauthorized, "UNAUTHORIZED", "invalid token")
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			utils.JSONError(c, http.StatusUnauthorized, "UNAUTHORIZED", "invalid token claims")
			c.Abort()
			return
		}

		// get sub (user id)
		sub, ok := claims["sub"]
		if !ok {
			utils.JSONError(c, http.StatusUnauthorized, "UNAUTHORIZED", "invalid token claims")
			c.Abort()
			return
		}

		// jwt uses float64 for numbers by default
		var uid uint
		switch v := sub.(type) {
		case float64:
			uid = uint(v)
		case int:
			uid = uint(v)
		case uint:
			uid = v
		default:
			utils.JSONError(c, http.StatusUnauthorized, "UNAUTHORIZED", "invalid subject claim")
			c.Abort()
			return
		}

		user, err := repositories.UserRepo.FindByID(uid)
		if err != nil {
			utils.JSONError(c, http.StatusUnauthorized, "UNAUTHORIZED", "user not found")
			c.Abort()
			return
		}

		// set user on context for handlers
		c.Set("currentUser", user)
		c.Next()
	}
}

func RequireAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		u, exists := c.Get("currentUser")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":  "UNAUTHORIZED",
				"detail": "user not authenticated",
			})
			return
		}

		user := u.(*models.User)

		if user.Role != "admin" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error":  "FORBIDDEN",
				"detail": "admin access required",
			})
			return
		}

		c.Next()
	}
}
