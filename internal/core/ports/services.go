package ports

import "github.com/thalesb16/koifood/internal/core/domain"

type StoreService interface {
	Get(id string) (domain.Store, error)
	Create(name string) (domain.Store, error)
}

type OrderService interface {
	Create(storeID, order string) (domain.Order, error)
}
