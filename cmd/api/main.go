package main

import (
	"crud-app/internal/config"
	"log"
	"net/http"
	"os"
)

func main() {

	router := config.SetupRouter()

	log.Println("Starting server on port 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Could not listen on port 8080 %v", err)
		os.Exit(1)
	}

}