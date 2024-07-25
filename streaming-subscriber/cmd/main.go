package main

import (
	"context"
	"fmt"
	daprd "github.com/dapr/go-sdk/client"
	"log"
)

func main() {
	client, err := daprd.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	if err := subscribeToOrders(client); err != nil {
		log.Fatal(err)
	}
}

func subscribeToOrders(client daprd.Client) error {
	deadLetterTopic := "deadletter"
	sub, err := client.Subscribe(context.Background(), daprd.SubscriptionOptions{
		PubsubName:      "pubsub",
		Topic:           "orders",
		DeadLetterTopic: &deadLetterTopic,
	})
	if err != nil {
		return err
	}
	defer sub.Close()
	fmt.Printf("Subscription created\n")

	for {
		msg, err := sub.Receive()
		if err != nil {
			return fmt.Errorf("error receiving message: %v", err)
		}

		log.Printf("Streaming Subscriber received: : %s\n", msg.RawData)
		if err := msg.Success(); err != nil {
			return fmt.Errorf("error sending message success: %v", err)
		}
	}
}
