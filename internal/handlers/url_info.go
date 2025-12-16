package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"url-shortener/internal/services"
)

type URLInfoHandler struct {
	service *services.URLService
}

func NewURLInfoHandler(s *services.URLService) *URLInfoHandler {
	return &URLInfoHandler{service: s}
}

func (h *URLInfoHandler) GetInfo(c *gin.Context) {
	codeURL := c.Param("codeURL") 

	url, err := h.service.GetInfo(codeURL)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "short url not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"shortURL":    url.ShortURL,
		"originalURL": url.OriginalURL,
		"clicks":      url.Clicks,
		"created_at":  url.CreatedAt,
	})
}
