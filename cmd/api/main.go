package main

import (
	"crud-app/internal/server"
	"net/http"
)

func main() {
	router := server.TestHandler()
	http.ListenAndServe(":8080", router)
}