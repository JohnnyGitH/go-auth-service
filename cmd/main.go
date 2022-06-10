package main

import (
	"log"
	"workspace/internal/server"
)

func init() {
	log.Println("this is init function")
}

// Main entry for the service
func main() {
	log.Println("this is the main function")
	// Get Config

	// Set up database
	// User database

	// Set up services
	// OAth service

	// Start Server
	server := server.NewHTTPServer()
	server.StartServer()
}
