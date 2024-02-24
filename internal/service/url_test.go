package service

import (
	"strings"
	"testing"
)

func TestStoreURL(t *testing.T) {
	service := NewUrlService()
	shortUrl, err := service.StoreShortURL("https://example.com")
	if err != nil {
		t.Fatalf("Erro ao armazenar Url: %v", err)
	}
	if !strings.HasPrefix(shortUrl, "www.") || !strings.HasSuffix(shortUrl, ".com") {
		t.Errorf("URL curta inválida: %v", shortUrl)
	}
}

func TestGetShortURL(t *testing.T){
	service := NewUrlService()
	originalUrl := "https://exemple.com"
	shortUrl, _ := service.StoreShortURL(originalUrl)
	retrieveUrl, exists := service.GetShortURL(shortUrl)
	if !exists {
		t.Errorf("URL não encontrada: %v", shortUrl)
	}
	if retrieveUrl != originalUrl {
		t.Errorf("URL recuperada não corresponde a URL original: esperado %v, obtido %v", originalUrl, retrieveUrl)
	}
}