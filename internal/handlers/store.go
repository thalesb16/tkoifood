package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thalesb16/tkoifood/internal/core/ports"
)

type StoreHandler struct {
	service ports.StoreService
}

func NewStoreHandler(service ports.StoreService) *StoreHandler {
	return &StoreHandler{
		service: service,
	}
}

func (h *StoreHandler) Create(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]bool{"ok": true})
}
