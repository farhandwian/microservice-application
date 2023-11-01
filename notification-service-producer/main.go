package main

import (
	"encoding/json"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type MessageType string

const (
	TypeEmail    MessageType = "email"
	TypeFirebase MessageType = "firebase"
)

type MessageContent interface{}

type EmailContent struct {
	Message string `json:"message"`
	To      string `json:"to"`
}

type FirebaseContent struct {
	Message string `json:"message"`
	To      string `json:"to"`
	Token   string `json:"token"`
}

type Message struct {
	Type    MessageType    `json:"type"`
	Content MessageContent `json:"content"`
}

func produce(p *kafka.Producer) {
	defer p.Close()

	endTime := time.Now().Add(15 * time.Second)
	content := &EmailContent{
		Message: "anda berhasil membeli produk saos abc seharga 15000",
		To:      "dwyanfarhan123@gmail.com",
	}

	firebase_content := &FirebaseContent{
		Message: "anda berhasil membeli produk saos abc seharga 15000",
		To:      "dwyanfarhan123@gmail.com",
		Token:   "fd8BwBI9rn8:APA91bGsu-7PGeZ7s-Xzbdhs0xxm-hOo_woJy6mWtYfVoby6nuep1Ll1smqelTbBihOubMPzpVTt3tXNmexYqveKWqCyL-M0N5PSlgyBV4KsSWJOonMWQpISllpt-PAq32zofA8pZ-Ji",
	}

	value := &Message{
		Type:    TypeEmail,
		Content: content,
	}

	firebase_value := &Message{
		Type:    TypeFirebase,
		Content: firebase_content,
	}

	alternate := true // Start with this flag

	for time.Now().Before(endTime) {
		var currentValue *Message
		if alternate {
			currentValue = value
		} else {
			currentValue = firebase_value
		}

		jsonData, err := json.Marshal(currentValue)
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
		alternate = !alternate // Toggle the flag for next iteration
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
