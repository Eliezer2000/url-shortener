package controller

import (
	"bytes"
	"crud-app/internal/service"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStoreShortUrlHandler(t *testing.T) {
	service := service.NewUrlService()
	controller := NewUrlController(service)

	t.Run("valid request", func (t *testing.T)  {
		requestBody := bytes.NewBuffer([]byte( `{"url": "https://example.com"}`))
		req, _ := http.NewRequest(http.MethodPost, "/url", requestBody)
		resp := httptest.NewRecorder()
		controller.StoreShortUrlHandler(resp, req)
		assert.Equal(t, http.StatusCreated, resp.Code)
		assert.NotEmpty(t, resp.Body.String())
	})

	t.Run("nil body", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPost, "/url", nil)
		resp := httptest.NewRecorder()
		controller.StoreShortUrlHandler(resp, req)
		assert.Equal(t, http.StatusBadRequest, resp.Code)
	})

	t.Run("Invalid url", func(t *testing.T) {
		requestBody := bytes.NewBuffer([]byte(`{"url": "invalid_url"}`))
		req, _ := http.NewRequest(http.MethodPost, "/url", requestBody)
		resp := httptest.NewRecorder()
		controller.StoreShortUrlHandler(resp, req)
		assert.Equal(t, http.StatusBadRequest, resp.Code)
	})

	t.Run("Unmarshal error", func(t *testing.T) {
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
	req, _ := http.NewRequest(http.MethodGet, "/url/shortUrl", nil)
	resp := httptest.NewRecorder()
	controller.GetShortURLHandler(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, "success", resp.Body.String())
}

func TestGenerateRandomURL(t *testing.T){
	result := generateRandomURL()
	assert.NotEmpty(t, result)
} 