package main

import (
    "github.com/gin-gonic/gin"

    "url-shortener/internal/database"
    "url-shortener/internal/handlers"
    "url-shortener/internal/services"
)

func main() {
    database.ConnectMongo()

    r := gin.Default()

    urlService := services.NewURLService()

    shortHandler := handlers.NewShortURLHandler(urlService)
    redirectHandler := handlers.NewRedirectHandler(urlService)

	// test short url 
    r.POST("/api/shorten", shortHandler.Create)

	// get new url
    r.GET("/:code", redirectHandler.Redirect)

	// PORT
    r.Run(":8080")
}
