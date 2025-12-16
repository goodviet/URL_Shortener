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
	codeURL := c.Param("codeURL")

	url, err := h.service.FindAndIncreaseClick(c.Request.Context(), codeURL)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	// Redirect 302
	c.Redirect(http.StatusFound, url.OriginalURL)
}
