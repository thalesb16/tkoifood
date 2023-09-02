package main

import (
	"github.com/thalesb16/tkoifood/internal/application/api"
	"github.com/thalesb16/tkoifood/internal/core/services"
	"github.com/thalesb16/tkoifood/internal/handlers"
	"github.com/thalesb16/tkoifood/internal/repositories"
	"github.com/thalesb16/tkoifood/pkg/uidgen"
)

func main() {
	storeRepository := repositories.NewStorage()
	storeService := services.NewStoreService(storeRepository, uidgen.New())
	storeHandler := handlers.NewStoreHandler(storeService)

	API := api.New(*storeHandler)

	if err := API.Run(); err != nil {
		panic(err)
	}
}
