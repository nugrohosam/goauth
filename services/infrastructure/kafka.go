package infrastructure

import (
	"context"

	"github.com/segmentio/kafka-go"
)

// KafkaConfig ..
type KafkaConfig struct {
	Writer *kafka.Writer
}

// ConfigureKafka ..
func (kafkaConfig *KafkaConfig) ConfigureKafka(url string, topic string) {
	kafkaConfig.Writer = &kafka.Writer{
		Addr:     kafka.TCP(url),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

// CloseKafkaConn ..
func (kafkaConfig *KafkaConfig) CloseKafkaConn() error {
	if err := kafkaConfig.Writer.Close(); err != nil {
		return err
	}

	return nil
}

// PushMessageKafka ..
func (kafkaConfig *KafkaConfig) PushMessageKafka(key, value string) error {
	err := kafkaConfig.Writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("Key-A"),
			Value: []byte("Hello World!"),
		},
	)

	if err != nil {
		return err
	}

	return nil
}
