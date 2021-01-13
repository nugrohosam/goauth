package kafka

import (
	"testing"

	"github.com/segmentio/kafka-go"
)

// KafkaTopicsTestRun ..
func KafkaTopicsTestRun(t *testing.T) {
	conn, err := kafka.Dial("tcp", "localhost:9092")
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	partitions, err := conn.ReadPartitions()
	if err != nil {
		panic(err.Error())
	}

	m := map[string]struct{}{}

	for _, p := range partitions {
		m[p.Topic] = struct{}{}
	}
	
	t.Log(m)
}
