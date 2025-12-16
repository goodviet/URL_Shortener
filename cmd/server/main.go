package main

import (
    "context"
    "log"
    "time"

    "github.com/gin-gonic/gin"
    "url-shortener/internal/database"
    "url-shortener/internal/models"
)

func main() {
    database.ConnectMongo()

    url := models.URL{
        ShortURL:      "test123",
        OriginalURL:  "https://example.com",
        Clicks:    0,
        CreatedAt: time.Now(),
    }

    _, err := database.URLCollection().InsertOne(context.Background(), url)
    if err != nil {
        log.Fatal(err)
    }

    log.Println("Inserted test document")

    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "pong"})
    })
    r.Run(":8080")
}
