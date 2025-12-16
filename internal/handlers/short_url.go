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

type CreateShortURLRequest struct {
    URL string `json:"url" binding:"required,url"`
}

func (h *ShortURLHandler) Create(c *gin.Context) {
    var req CreateShortURLRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    url, err := h.service.CreateShortURL(
        c.Request.Context(),
        req.URL,
    )
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }

    c.JSON(200, gin.H{
        "short_code": url.ShortURL,
        "short_url":  "http://localhost:8080/" + url.ShortURL,
    })
}
