package services

import (
    "context"
    "time"

    "url-shortener/internal/database"
    "url-shortener/internal/models"
)

type URLService struct{}

func NewURLService() *URLService {
    return &URLService{}
}

func (s *URLService) CreateShortURL(originalURL string, shortURL string) error {
    url := models.URL{
        OriginalURL: originalURL,
        ShortURL:    shortURL,
        Clicks:      0,
        CreatedAt:   time.Now(),
    }

    _, err := database.URLCollection().InsertOne(context.Background(), url)
    return err
}

func (s *URLService) FindByShortURL(codeURL string) (*models.URL, error) {
    var url models.URL
    err := database.URLCollection().
        FindOne(context.Background(), map[string]string{"short_url": codeURL}).
        Decode(&url)

    return &url, err
}
