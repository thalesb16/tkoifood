package services

import (
	"github.com/thalesb16/tkoifood/internal/core/domain"
	"github.com/thalesb16/tkoifood/internal/core/ports"
	"github.com/thalesb16/tkoifood/pkg/uidgen"
)

type StoreService struct {
	repository ports.StoreRepository
	uidGen     uidgen.UIDGen
}

func NewStoreService(storeRepository ports.StoreRepository, uidGen uidgen.UIDGen) *StoreService {
	return &StoreService{
		repository: storeRepository,
		uidGen:     uidGen,
	}
}

func (srv *StoreService) Get(id string) (domain.Store, error) {
	store, err := srv.repository.Get(id)
	if err != nil {
		return domain.Store{}, err
	}

	return store, nil
}

func (srv *StoreService) Create(name string) (string, error) {
	newStore := domain.NewStore(srv.uidGen.New(), name)

	err := srv.repository.Save(newStore)
	if err != nil {
		return "", err
	}

	return newStore.ID, nil
}
