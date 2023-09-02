package services

import (
	"github.com/thalesb16/koifood/internal/core/domain"
	"github.com/thalesb16/koifood/internal/core/ports"
	"github.com/thalesb16/koifood/pkg/uidgen"
)

type OrderService struct {
	queue  ports.OrderQueue
	uidGen uidgen.UIDGen
}

func NewOrderService(orderQueue ports.OrderQueue, uidGen uidgen.UIDGen) *OrderService {
	return &OrderService{
		queue:  orderQueue,
		uidGen: uidGen,
	}
}

func (srv *OrderService) Create(order string) (domain.Order, error)
