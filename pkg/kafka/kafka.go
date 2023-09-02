package kafka

import (
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

var topic string = "tkoifood-orders"
var server string = "localhost:19092"
var group string = "consumer-group"

type Kafka interface {
	Read(duration time.Duration) ([]byte, error)
	Write(msg []string, headers []map[string]string) error
}

type client struct {
	*kafka.Consumer
	*kafka.Producer
}

func NewConsumer() (Kafka, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": server,
		"group.id":          group,
	})
	if err != nil {
		return nil, err
	}

	c.SubscribeTopics([]string{topic}, nil)

	return &client{c, nil}, nil
}

func NewProducer() (Kafka, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": server,
	})
	if err != nil {
		return nil, err
	}

	return &client{nil, p}, nil
}

func (c client) Read(duration time.Duration) ([]byte, error) {
	msg, err := c.ReadMessage(duration)
	return msg.Value, err
}

func (c client) Write(msg []string, headers []map[string]string) error {
	newHeaders := []kafka.Header{}
	for _, h := range headers {
		nH := kafka.Header{
			Key:   h["key"],
			Value: []byte(h["value"]),
		}

		newHeaders = append(newHeaders, nH)
	}

	for _, w := range msg {
		err := c.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Headers:        newHeaders,
			Key:            newHeaders[0].Value,
			Value:          []byte(w),
		}, nil)
		if err != nil {
			return err
		}
	}

	c.Flush(15 * 100)
	return nil
}
