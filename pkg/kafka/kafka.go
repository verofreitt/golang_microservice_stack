package kafka

import (
	"context"
	"golang_microservice_stack/config"

	"github.com/segmentio/kafka-go"
)

func NewKafkaConn(cfg *config.Config) (*kafka.Conn, error) {
	return kafka.DialContext(context.Background(), "tcp", cfg.Kafka.Brokers[0])
}
