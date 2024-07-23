package main

import (
	"context"
	"errors"
	"fmt"
	daprd "github.com/dapr/go-sdk/client"
	"log"
)

func main() {
	client, err := daprd.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	deadLetterTopic := "deadletter"
	sub, err := client.Subscribe(context.Background(), daprd.SubscriptionOptions{
		PubsubName:      "pubsub",
		Topic:           "orders",
		DeadLetterTopic: &deadLetterTopic,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(">>Created subscription\n")

	for {
		msg, err := sub.Receive()
		if err != nil {
			log.Fatalf("error receiving message: %v", err)
		}
		log.Printf("Streaming Subscriber received: : %s\n", msg.RawData)
		if err := msg.Success(); err != nil {
			log.Fatalf("error sending message success: %v", err)
		}
	}

	if err := errors.Join(sub.Close()); err != nil {
		log.Fatal(err)
	}
}
