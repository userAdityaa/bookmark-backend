package main

import (
	"log"
	"net/http"

	"github.com/userAdityaa/bookmark-backend/config"
	"github.com/userAdityaa/bookmark-backend/routes"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer config.DisconnectDB(db)
	router := routes.SetUpRoutes(db)
	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal("Failed to start the server:", err)
	}
}
