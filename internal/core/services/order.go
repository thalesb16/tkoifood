package services

import (
	"github.com/thalesb16/tkoifood/internal/core/domain"
	"github.com/thalesb16/tkoifood/internal/core/ports"
	"github.com/thalesb16/tkoifood/pkg/uidgen"
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

func (srv *OrderService) Create(storeID, order string) (string, error) {
	newOrder := domain.NewOrder(srv.uidGen.New(), storeID, order)

	err := srv.queue.Write(newOrder)
	if err != nil {
		return "", err
	}

	return newOrder.ID, err
}
