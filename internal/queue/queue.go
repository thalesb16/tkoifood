package queue

import (
	"github.com/thalesb16/tkoifood/internal/core/domain"
	"github.com/thalesb16/tkoifood/pkg/kafka"
)

type queue struct {
	client kafka.Kafka
}

func NewQueue(client kafka.Kafka) *queue {
	return &queue{
		client: client,
	}
}

func (q *queue) Write(order domain.Order) error {
	msg := []string{order.Order}
	key := "StoreID"
	value := order.StoreID
	headers := []map[string]string{
		{
			"key":   key,
			"value": value,
		},
	}
	return q.client.Write(msg, headers)
}
