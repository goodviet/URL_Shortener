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
	infoHandler := handlers.NewURLInfoHandler(urlService)
	// get list url tao tu url original
	listHandler := handlers.NewListURLHandler(urlService)

    r.POST("/api/shorten", shortHandler.Create)
    r.GET("/:codeURL", redirectHandler.Redirect)
	r.GET("/api/links/:codeURL", infoHandler.GetInfo)
	r.GET("/api/links", listHandler.List)

    r.Run(":8080")
}
