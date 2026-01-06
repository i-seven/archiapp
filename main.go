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

		// Allow specific origins
		allowedOrigins := []string{
			"http://localhost:5173",
			"http://127.0.0.1:5173",
			"http://localhost:3000",
			"http://localhost:8080",
			"*", // Allow all for testing
		}

		for _, allowedOrigin := range allowedOrigins {
			if allowedOrigin == "*" || origin == allowedOrigin {
				c.Writer.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
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
	r.Static("/images", config.ImageDir)

	// Use CORS middleware FIRST
	r.Use(CORS())

	// Logger middleware
	r.Use(func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		c.Next()

		latency := time.Since(start)
		fmt.Printf("[%s] %s %s | Status: %d | Latency: %v\n",
			method, path, c.ClientIP(), c.Writer.Status(), latency)
	})

	// Test endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Register ALL routes - ADD THIS LINE
	routes.RegisterUserRoutes(r)
	routes.RegisterProductRoutes(r) // ← THIS WAS MISSING!

	// Handle 404 with JSON (not HTML)
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"error": "Route not found",
			"path":  c.Request.URL.Path,
		})
	})

	// Print registered routes
	fmt.Println("\n=== Registered Routes ===")
	for _, route := range r.Routes() {
		fmt.Printf("%-6s %s\n", route.Method, route.Path)
	}
	fmt.Println("=======================")

	// Start server
	println("Server starting on :8080...")
	r.Run(":8080")
}
