package repositories

import "github.com/thalesb16/koifood/internal/core/domain"

type storage struct {
	stores map[string][]byte
}

func NewStorage() *storage {
	return &storage{stores: map[string][]byte{}}
}

func (repo *storage) Get(id string) (domain.Store, error) { return domain.Store{}, nil }
func (repo *storage) Save(store domain.Store) error       { return nil }
