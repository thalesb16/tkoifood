package ports

import "github.com/thalesb16/tkoifood/internal/core/domain"

type StoreRepository interface {
	Get(id string) (domain.Store, error)
	Save(store domain.Store) error
}
