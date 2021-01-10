package main

import (
	"context"
	"fmt"

	infrastructure "github.com/nugrohosam/gosampleapi/services/infrastructure"
)

// main ...
func main() {
	writer, err := infrastructure.ConfigureKafka(
		[]string{
			"localhost:9092",
		}, "1", "test",
	)

	err = infrastructure.KafkaPushMessage(context.Background(), nil, []byte("your message content"))
	infrastructure.KafkaProducerClose(writer)

	fmt.Println(err.Error())
}
