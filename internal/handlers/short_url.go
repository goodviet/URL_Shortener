package handlers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "url-shortener/internal/services"
)

type ShortURLHandler struct {
    service *services.URLService
}

func NewShortURLHandler(s *services.URLService) *ShortURLHandler {
    return &ShortURLHandler{service: s}
}

func (h *ShortURLHandler) Create(c *gin.Context) {
    var req struct {
        URL string `json:"url"`
    }

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    err := h.service.CreateShortURL(req.URL, "abc123")
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }

    c.JSON(200, gin.H{"short_url": "http://localhost:8080/abc123"})
}
