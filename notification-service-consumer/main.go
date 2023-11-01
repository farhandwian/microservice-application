package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/smtp"
	"sync"

	firebase "firebase.google.com/go"
	"github.com/confluentinc/confluent-kafka-go/kafka"

	// "firebase.google.com/go/v4/messaging"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
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

func sendEmail(msg EmailContent) {
	go func(msg EmailContent) {
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

func SendPushNotification(deviceTokens string) error {
	opt := option.WithCredentialsFile("serviceAccountKey.json")

	app, err := firebase.NewApp(context.Background(), nil, opt)

	if err != nil {
		log.Printf("Error in initializing firebase : %s", err)
		return err
	}

	fcmClient, err := app.Messaging(context.Background())

	if err != nil {
		return err
	}

	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: "Congratulations!!",
			Body:  "You have just implemented push notification",
		},
		Token: deviceTokens,
	}

	_, err = fcmClient.Send(context.Background(), message)
	if err != nil {
		return err
	}

	log.Println("Push notification sent to single device!")

	return nil
}

func (m *Message) UnmarshalJSON(data []byte) error {
	type Alias Message

	aux := &struct {
		Content json.RawMessage `json:"content"`
		*Alias
	}{
		Alias: (*Alias)(m),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	switch m.Type {
	case TypeEmail:
		var emailContent EmailContent
		if err := json.Unmarshal(aux.Content, &emailContent); err != nil {
			return err
		}
		m.Content = emailContent
	case TypeFirebase:
		var firebaseContent FirebaseContent
		if err := json.Unmarshal(aux.Content, &firebaseContent); err != nil {
			return err
		}
		m.Content = firebaseContent
	}
	return nil
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

		switch decode_msg.Type {
		case TypeEmail:
			if content, valid := decode_msg.Content.(EmailContent); valid {
				sendEmail(content)
			} else {
				fmt.Println("error decode message")
			}
		case TypeFirebase:
			if content, valid := decode_msg.Content.(FirebaseContent); valid {
				SendPushNotification(content.Token) // adjust the function to handle a single token or change the way you pass tokens
			} else {
				fmt.Println("error decode message")
			}
		default:
			fmt.Printf("Unknown message type: %s\n", decode_msg.Type)
		}
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
