package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"url-shortener/internal/services"
)

type ListURLHandler struct {
	service *services.URLService
}

func NewListURLHandler(s *services.URLService) *ListURLHandler {
	return &ListURLHandler{service: s}
}

func (h *ListURLHandler) List(c *gin.Context) {
	urls, err := h.service.ListAll()
	if err != nil {
		c.JSON(500, gin.H{"error": "cannot list urls"})
		return
	}

	c.JSON(http.StatusOK, urls)
}
