package ports

import "github.com/thalesb16/tkoifood/internal/core/domain"

type StoreService interface {
	Get(id string) (domain.Store, error)
	Create(name string) (string, error)
}

type OrderService interface {
	Create(storeID, order string) (domain.Order, error)
}
