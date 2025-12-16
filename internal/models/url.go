package models

import (
    "time"

    "go.mongodb.org/mongo-driver/bson/primitive"
)

type URL struct {
    ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    ShortURL      string             `bson:"shortURL" json:"shortURL"`
    OriginalURL  string             `bson:"originalURL" json:"originalURL"`
    Clicks    int                `bson:"clicks" json:"clicks"`
    CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}
