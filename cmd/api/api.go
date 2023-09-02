package main

import (
	"github.com/thalesb16/tkoifood/internal/application/api"
	"github.com/thalesb16/tkoifood/internal/core/services"
	"github.com/thalesb16/tkoifood/internal/handlers"
	"github.com/thalesb16/tkoifood/internal/queue"
	"github.com/thalesb16/tkoifood/internal/repositories"
	"github.com/thalesb16/tkoifood/pkg/kafka"
	"github.com/thalesb16/tkoifood/pkg/uidgen"
)

func main() {
	storeRepository := repositories.NewStorage()
	storeService := services.NewStoreService(storeRepository, uidgen.New())
	storeHandler := handlers.NewStoreHandler(storeService)

	kafkaClient, err := kafka.NewProducer()
	if err != nil {
		panic(err)
	}
	queue := queue.NewQueue(kafkaClient)
	orderService := services.NewOrderService(queue, uidgen.New())
	orderHandler := handlers.NewOrderHandler(orderService)

	API := api.New(*storeHandler, *orderHandler)

	if err := API.Run(); err != nil {
		panic(err)
	}
}
