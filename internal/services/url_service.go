package services

import (
	"context"
	"errors"
	"time"

	"url-shortener/internal/database"
	"url-shortener/internal/models"
	"url-shortener/internal/utils"

	"go.mongodb.org/mongo-driver/bson"
)

type URLService struct{}

func NewURLService() *URLService {
	return &URLService{}
}

// Tạo short URL
func (s *URLService) CreateShortURL(ctx context.Context, originalURL string) (*models.URL, error) {
	collection := database.URLCollection()

	var shortURL string
	for {
		shortURL = utils.GenerateShortUrl(6)
		exists, err := s.checkShortURLExists(ctx, shortURL)
		if err != nil {
			return nil, err
		}
		if !exists {
			break
		}
	}

	url := models.URL{
		ShortURL:     shortURL,
		OriginalURL: originalURL,
		Clicks:      0,
		CreatedAt:   time.Now(),
	}

	_, err := collection.InsertOne(ctx, url)
	if err != nil {
		return nil, errors.New("cannot insert short url")
	}

	return &url, nil
}

// Redirect + tăng click
func (s *URLService) FindAndIncreaseClick(ctx context.Context, shortURL string) (*models.URL, error) {
	var url models.URL

	err := database.URLCollection().FindOneAndUpdate(
		ctx,
		bson.M{"shortURL": shortURL},
		bson.M{"$inc": bson.M{"clicks": 1}},
	).Decode(&url)

	if err != nil {
		return nil, err
	}

	return &url, nil
}

// Check trùng short code
func (s *URLService) checkShortURLExists(ctx context.Context, shortURL string) (bool, error) {
	count, err := database.URLCollection().
		CountDocuments(ctx, bson.M{"shortURL": shortURL})
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// get info url

func (s *URLService) GetInfo(codeURL string) (*models.URL, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var url models.URL
	err := database.URLCollection().
		FindOne(ctx, bson.M{"shortURL": codeURL}).
		Decode(&url)

	if err != nil {
		return nil, err
	}

	return &url, nil
}


// 
func (s *URLService) ListAll() ([]models.URL, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := database.URLCollection().Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var urls []models.URL
	if err := cursor.All(ctx, &urls); err != nil {
		return nil, err
	}

	return urls, nil
}
