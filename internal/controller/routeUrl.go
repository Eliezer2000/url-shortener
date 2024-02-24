package controller

import (
	"crud-app/internal/service"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type URLController struct {
	service service.UrlService
}

type URLRequest struct {
	Url string `json:"url"`
}


func NewUrlController(service service.UrlService) *URLController {
	return &URLController {
		service: service,
	}
}

func (urlController URLController) StoreShortUrlHandler(w http.ResponseWriter, r *http.Request)  {
	if r.Body == nil {
		http.Error(w, "request body is nil", http.StatusBadRequest)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading request body: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var urlRequest URLRequest
	err = json.Unmarshal(body, &urlRequest)
	if err != nil {
		log.Printf("Error unmarshalling request body: %v ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	shortUrl, err := urlController.service.StoreShortURL(urlRequest.Url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(shortUrl))
}

func (urlController URLController) GetShortURLHandler(w http.ResponseWriter, r *http.Request){
	shortUrl := r.URL.Path[len("/url/"):]
	originalUrl, exists := urlController.service.GetShortURL(shortUrl)
	if !exists {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(originalUrl))
}

