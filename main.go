package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/example.com/cars/api"
	"github.com/example.com/cars/db"
	"github.com/gorilla/mux"
)

func main() {
	// Initialize the database
	db.InitDB()

	// Initialize Gorilla Mux router
	r := mux.NewRouter()

	// Initialize routes
	api.InitCarRoutes(r)
	// Add more route initializations if needed

	// Start the server
	port := ":8080"
	fmt.Printf("Server is running on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}
