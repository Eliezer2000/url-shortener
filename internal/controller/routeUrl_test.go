package controller

import (
	"bytes"
	"crud-app/internal/service"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStoreShortUrlHandler(t *testing.T) {
	service := service.NewUrlService()
	controller := NewUrlController(service)

	t.Run("valid request with URL 'https://example.com' should return status 201 and non-empty body", func (t *testing.T)  {
		requestBody := bytes.NewBuffer([]byte( `{"url": "https://example.com"}`))
		req, _ := http.NewRequest(http.MethodPost, "/url", requestBody)
		resp := httptest.NewRecorder()
		controller.StoreShortUrlHandler(resp, req)
		assert.Equal(t, http.StatusCreated, resp.Code)
		assert.NotEmpty(t, resp.Body.String())
	})

	t.Run("request with nil body should return status 400", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPost, "/url", nil)
		resp := httptest.NewRecorder()
		controller.StoreShortUrlHandler(resp, req)
		assert.Equal(t, http.StatusBadRequest, resp.Code)
	})

	t.Run("Request with invalid URL 'invalid_url' should return status 400", func(t *testing.T) {
		requestBody := bytes.NewBuffer([]byte(`{"url": "invalid_url"}`))
		req, _ := http.NewRequest(http.MethodPost, "/url", requestBody)
		resp := httptest.NewRecorder()
		controller.StoreShortUrlHandler(resp, req)
		assert.Equal(t, http.StatusBadRequest, resp.Code)
	})

	t.Run("Request with unmarshalable body '{\"url\":123}' should return status 400", func(t *testing.T) {
		requestBody := bytes.NewBuffer([]byte(`{"url": 123}`))
		req, _ := http.NewRequest(http.MethodPost, "/url", requestBody)
		resp := httptest.NewRecorder()
		controller.StoreShortUrlHandler(resp, req)
		assert.Equal(t, http.StatusBadRequest, resp.Code)
	})
}


func TestGetShortUrlHandler(t *testing.T) {
	service := service.NewUrlService()
	controller := NewUrlController(service)
	shortUrl, err := service.StoreShortURL("https://www.zasdeki.com")
	if err != nil {
		log.Fatalf("Falha ao armazenar URL: %v", err)
	}

	t.Run("Request to '/url/"+shortUrl+"should return status 200 and body 'success'", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/url/"+shortUrl, nil)
		resp := httptest.NewRecorder()
		controller.GetShortURLHandler(resp, req)
		assert.Equal(t, http.StatusOK, resp.Code)
		assert.Equal(t, "https://www.zasdeki.com", resp.Body.String())
	})

	t.Run("Request to '/url/noneexistent' should return status 404", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/url/noneexistent", nil)
		resp := httptest.NewRecorder()
		controller.GetShortURLHandler(resp, req)
		assert.Equal(t, http.StatusNotFound, resp.Code)
		assert.Equal(t, "URL not found\n", resp.Body.String())
	})

}

