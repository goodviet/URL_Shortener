package services

import (
    "context"
    "time"

    "url-shortener/internal/database"
    "url-shortener/internal/models"

    "go.mongodb.org/mongo-driver/bson"
)

type URLService struct {
}

func NewURLService() *URLService {
    return &URLService{}
}

func (s *URLService) Create(originalURL string, shortURL string) error {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    url := models.URL{
        ShortURL:      shortURL,
        OriginalURL:  originalURL,
        Clicks:    0,
        CreatedAt: time.Now(),
    }

    _, err := database.URLCollection().InsertOne(ctx, url)
    return err
}

func (s *URLService) FindByCode(shortURL string) (*models.URL, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    var url models.URL
    err := database.URLCollection().
        FindOne(ctx, bson.M{"shortURL": shortURL}).
        Decode(&url)

    if err != nil {
        return nil, err
    }

    return &url, nil
}

func (s *URLService) IncreaseClick(shortURL string) error {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    _, err := database.URLCollection().UpdateOne(
        ctx,
        bson.M{"shortURL": shortURL},
        bson.M{"$inc": bson.M{"clicks": 1}},
    )

    return err
}
