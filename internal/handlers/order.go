package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thalesb16/tkoifood/internal/core/domain"
	"github.com/thalesb16/tkoifood/internal/core/ports"
)

type OrderHandler struct {
	service ports.OrderService
}

type OrderDTO domain.Order

func NewOrderHandler(service ports.OrderService) *OrderHandler {
	return &OrderHandler{
		service: service,
	}
}

func (h *OrderHandler) Create(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	order := OrderDTO{}
	err = json.Unmarshal(body, &order)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	orderID, err := h.service.Create(order.StoreID, order.Order)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	order.ID = orderID

	c.JSON(http.StatusOK, order)
}
