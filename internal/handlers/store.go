package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thalesb16/tkoifood/internal/core/domain"
	"github.com/thalesb16/tkoifood/internal/core/ports"
)

type StoreHandler struct {
	service ports.StoreService
}

type StoreDTO domain.Store

func NewStoreHandler(service ports.StoreService) *StoreHandler {
	return &StoreHandler{
		service: service,
	}
}

func (h *StoreHandler) Create(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	store := StoreDTO{}
	err = json.Unmarshal(body, &store)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	storeID, err := h.service.Create(store.Name)
	err = json.Unmarshal(body, &store)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	store.ID = storeID

	c.JSON(http.StatusOK, store)
}

func (h *StoreHandler) Get(c *gin.Context) {
	storeID := c.Param("store_id")
	store, err := h.service.Get(storeID)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, store)
}
