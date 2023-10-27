package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/smtp"
	"sync"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Message struct {
	Typee   string `json:"type"`
	Message string `json:"message"`
	To      string `json:"email"`
}

func sendEmail(msg Message) {
	go func(msg Message) {
		from := "dwyanfarhan@gmail.com"
		password := "yvgogxiooxpbqrdo"
		to := msg.To

		body := "Order received: " + fmt.Sprintf("%+v", msg)
		send_msg := "From: " + from + "\n" +
			"To: " + to + "\n" +
			"Subject: Order Notification\n\n" +
			body

		err := smtp.SendMail("smtp.gmail.com:587",
			smtp.PlainAuth("", from, password, "smtp.gmail.com"),
			from, []string{to}, []byte(send_msg))

		if err != nil {
			log.Printf("Failed to send email: %v", err)
			return
		}

		fmt.Println("Email sent!")
	}(msg)
}

func startConsumer(instanceID int, wg *sync.WaitGroup) {
	defer wg.Done()
	config := &kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "test_group",
		"auto.offset.reset": "earliest",
	}

	consumer, err := kafka.NewConsumer(config)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	consumer.Subscribe(topic, nil)

	for {
		msg, err := consumer.ReadMessage(-1)
		if err != nil {
			fmt.Printf("Consumer %d: Read message with value %s from partition %d\n", instanceID, string(msg.Value), msg.TopicPartition.Partition)
			continue
		}

		var decode_msg Message
		if err := json.Unmarshal(msg.Value, &decode_msg); err != nil {
			fmt.Println("Error decoding message:", err)
			continue
		}

		fmt.Println(decode_msg)
		sendEmail(decode_msg)
	}
}

func main() {
	numConsumers := 3
	var wg sync.WaitGroup
	wg.Add(numConsumers)
	for i := 0; i < numConsumers; i++ {
		go startConsumer(i, &wg)
	}
	wg.Wait()
}

var topic = "notifications"
