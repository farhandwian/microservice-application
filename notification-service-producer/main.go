package main

import (
	"encoding/json"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Message struct {
	Typee   string `json:"type"`
	Message string `json:"message"`
	To      string `json:"email"`
}

func produce(p *kafka.Producer) {
	defer p.Close()

	endTime := time.Now().Add(15 * time.Second)
	value := &Message{
		Typee:   "order",
		Message: "anda berhasil membeli produk saos abc seharga 15000",
		To:      "dwyanfarhan123@gmail.com",
	}

	for time.Now().Before(endTime) {
		jsonData, err := json.Marshal(value)
		if err != nil {
			panic(err)
		}

		message := &kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &Topic,
				Partition: kafka.PartitionAny,
			},
			Value: jsonData,
		}

		p.Produce(message, nil)
		time.Sleep(1 * time.Second)
	}

	p.Flush(15 * 1000)
}

func main() {
	config := &kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
	}

	producer, err := kafka.NewProducer(config)
	if err != nil {
		panic(err)
	}
	produce(producer)
}

var Topic = "notifications"
