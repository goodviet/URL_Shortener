package handlers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "url-shortener/internal/services"
)

type RedirectHandler struct {
    service *services.URLService
}

func NewRedirectHandler(s *services.URLService) *RedirectHandler {
    return &RedirectHandler{service: s}
}

func (h *RedirectHandler) Redirect(c *gin.Context) {
    code := c.Param("code")

    url, err := h.service.FindByShortURL(code)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
        return
    }

    c.Redirect(http.StatusFound, url.OriginalURL)
}
