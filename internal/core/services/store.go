package services

import (
	"github.com/thalesb16/koifood/internal/core/domain"
	"github.com/thalesb16/koifood/internal/core/ports"
	"github.com/thalesb16/koifood/pkg/uidgen"
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

func (srv *StoreService) Get(id string) (domain.Store, error)

func (srv *StoreService) Create(name string) (domain.Store, error)
