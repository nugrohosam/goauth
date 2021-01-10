package kafka

import (
	"context"
	"log"

	"time"

	kafka "github.com/segmentio/kafka-go"
)

// Service ..
type Service struct {
	Conn *kafka.Conn
}

// NewConn ..
func (kafkaService *Service) NewConn(host, port, topic string, partition int) {
	var err error
	kafkaService.Conn, err = kafka.DialLeader(context.Background(), "tcp", host+":"+port, topic, partition)
	if err != nil {
		panic(err)
	}
}

// Produce ..
func (kafkaService *Service) Produce(partition int, message string) {
	var err error
	kafkaService.Conn.SetWriteDeadline(time.Now().Add(time.Second))
	_, err = kafkaService.Conn.WriteMessages(
		kafka.Message{Value: []byte(message)},
	)

	if err != nil {
		log.Fatal("Failed to write messages:", err)
	}

	if err = kafkaService.Conn.Close(); err != nil {
		log.Fatal("Failed to close writer:", err)
	}
}

// Consume ..
func (kafkaService *Service) Consume(partition int, topic string) string {

	kafkaService.Conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	batch := kafkaService.Conn.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max

	b := make([]byte, 10e3) // 10KB max per message
	for {
		_, err := batch.Read(b)
		if err != nil {
			break
		}
	}

	if err := batch.Close(); err != nil {
		panic(err)
	}

	if err := kafkaService.Conn.Close(); err != nil {
		panic(err)
	}

	return string(b)
}
