package main

import (
	"backendAf/config"
	"backendAf/routes"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnv()
	config.DbConInit()
	config.SyncDB()
}

// CORS Middleware
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")

		// Allow specific origins (you can change to "*" for all)
		allowedOrigins := []string{
			"http://localhost:5173", // Add this
			"http://127.0.0.1:5173", // Add this
		}

		for _, allowedOrigin := range allowedOrigins {
			if origin == allowedOrigin {
				c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
				break
			}
		}

		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		c.Next()

		latency := time.Since(start)
		fmt.Printf("[%s] %s %s | Status: %d | Latency: %v\n",
			method, path, c.ClientIP(), c.Writer.Status(), latency)
	})
	// Use CORS middleware
	r.Use(CORS())

	// Test endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Register routes
	routes.RegisterUserRoutes(r)

	// Start server
	println("Server starting on :8080...")
	r.Run(":8080")
}
