package main

import (
	"fmt"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "YOUR_KAFKA_BROKER_IP:9092",
		"acks":              "all",
		"security.protocol": "SASL_PLAINTEXT",
		"sasl.mechanisms":   "SCRAM-SHA-512",
		"sasl.username":     "ACCOUNT_NAME",
		"sasl.password":     "PASSWORD",
	})
	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Created Producer %v\n", p)

	topic := "test.testtopic.v0"

	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Successfully produced record to topic %s partition [%d] @ offset %v\n",
						*ev.TopicPartition.Topic, ev.TopicPartition.Partition, ev.TopicPartition.Offset)
				}
			}
		}
	}()
	totalMsgcnt := 10
	msgcnt := 0
	for msgcnt < totalMsgcnt {
		value := fmt.Sprintf("Producer example, message #%d", msgcnt)

		err = p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(value),
			Headers:        []kafka.Header{{Key: "myTestHeader", Value: []byte("header values are binary")}},
		}, nil)

		if err != nil {
			if err.(kafka.Error).Code() == kafka.ErrQueueFull {
				time.Sleep(time.Second)
				continue
			}
			fmt.Printf("Failed to produce message: %v\n", err)
		}
		msgcnt++
	}

	// Flush and close the producer and the events channel
	for p.Flush(10000) > 0 {
		fmt.Print("Still waiting to flush outstanding messages\n")
	}
	p.Close()
}
