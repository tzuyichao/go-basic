package main

import (
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {
	a, err := kafka.NewAdminClient(&kafka.ConfigMap{
		"bootstrap.servers": "YOUR_KAFKA_BROKER:9093",
		"acks":              "all",
		"security.protocol": "SASL_PLAINTEXT",
		"sasl.mechanisms":   "SCRAM-SHA-512",
		"sasl.username":     "ACCOUNT_NAME",
		"sasl.password":     "PASSWORD",
	})
	if err != nil {
		fmt.Printf("Failed to create Admin client: %s\n", err)
		os.Exit(1)
	}
	defer a.Close()

	m, err := a.GetMetadata(nil, true, 10000)
	if err != nil {
		fmt.Printf("Failed to get metadata with client-level error %s\n", err)
		os.Exit(1)
	}

	topics := m.Topics
	fmt.Printf("A total of %d topic(s) listed:\n", len(topics))
	for _, topic := range topics {
		fmt.Printf("Topic: %s\n", topic.Topic)
	}
}
