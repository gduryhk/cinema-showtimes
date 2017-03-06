package main

import (
	"log"
	"net/http"

	"github.com/gduryhk/cinema-showtimes/common"
	"github.com/gduryhk/cinema-showtimes/routers"
)

// Entry point for the program
func main() {

	// Calls startup logic
	common.StartUp()
	// Get the mux router object
	router := routers.InitRoutes()

	server := &http.Server{
		Addr:    common.AppConfig.Server,
		Handler: router,
	}
	log.Println("Listening...")
	server.ListenAndServe()
}
