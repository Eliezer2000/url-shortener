package controller

import (
	"crud-app/internal/service"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/url"
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
	_, err = url.ParseRequestURI(urlRequest.Url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	shortUrl := generateRandomURL()
	formatteUrl := fmt.Sprintf("www.%s.com", shortUrl)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(formatteUrl))
}

func (urlController URLController) GetShortURLHandler(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("success"))
}

func generateRandomURL() string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyz")
	b := make([]byte, 7)
	for i := range b {
		b[i] = byte(letters[rand.Intn(len(letters))])
	}
	return string(b)
}