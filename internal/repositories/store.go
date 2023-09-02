package repositories

import (
	"encoding/json"
	"fmt"

	"github.com/thalesb16/tkoifood/internal/core/domain"
)

type storage struct {
	stores map[string][]byte
}

func NewStorage() *storage {
	return &storage{stores: map[string][]byte{}}
}

func (repo *storage) Get(id string) (domain.Store, error) {
	if value, ok := repo.stores[id]; ok {
		store := domain.Store{}
		err := json.Unmarshal(value, &store)
		if err != nil {
			return domain.Store{}, err
		}

		return store, nil
	}

	return domain.Store{}, fmt.Errorf("store not found")
}

func (repo *storage) Save(store domain.Store) error {
	bytes, err := json.Marshal(store)
	if err != nil {
		return err
	}

	repo.stores[store.ID] = bytes
	return nil
}
