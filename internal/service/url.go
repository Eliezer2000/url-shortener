package service

import (
	"fmt"
	"math/rand"
	"net/url"
)

type UrlService struct {
	urlMap map[string]string
}

func NewUrlService() UrlService {
	return UrlService{
		urlMap: make(map[string]string),
	}
}

func (urlService *UrlService) StoreShortURL(urlRequest string) (string, error) {
	_, err := url.ParseRequestURI(urlRequest)
	if err != nil {
		return "", err
	}
	shortUrl := urlService.generateRandomURL()
	formattedUrl := fmt.Sprintf("www.%s.com", shortUrl)
	urlService.urlMap[formattedUrl] = urlRequest
	return formattedUrl,nil
}

func (urlService UrlService) GetShortURL (shortUrl string) (string, bool) {
	originalUrl, exists := urlService.urlMap[shortUrl]
	return originalUrl, exists
}

func(urlService UrlService) generateRandomURL() string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyz")
	b := make([]byte, 7)
	for i := range b {
		b[i] = byte(letters[rand.Intn(len(letters))])
	}
	return string(b)
}