package main

import (
	"backendAf/config"
	"backendAf/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnv()
	config.DbConInit()
	config.SyncDB()
}

func main() {
	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// 1. FIRST: Serve static images
	r.Static("/images", "./storage/images")

	// 2. SECOND: Register all routes
	routes.RegisterRoutes(r)

	//// 3. THIRD: Setup HTTPS
	// m := &autocert.Manager{
	// 	Cache:      autocert.DirCache("certs"),
	// 	Prompt:     autocert.AcceptTOS,
	// 	HostPolicy: autocert.HostWhitelist("example.com"), // Change to your domain
	// }

	// server := &http.Server{
	// 	Addr:    ":80",
	// 	Handler: r, // This r now has all routes registered
	// 	// TLSConfig: m.TLSConfig(),
	// }

	// 4. Start HTTP->HTTPS redirector
	// go http.ListenAndServe(":80", m.HTTPHandler(nil))

	// 5. FINALLY: Start the HTTPS server
	log.Printf("Starting HTTPS server on :80")
	// log.Fatal(server.ListenAndServe())
	r.Run(":80")
	// log.Fatal(server.ListenAndServeTLS("cert.pem", "key.pem"))
}
