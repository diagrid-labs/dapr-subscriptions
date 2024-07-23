package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	daprd "github.com/dapr/go-sdk/client"
	"github.com/dapr/go-sdk/service/common"
)

func main() {
	client, err := daprd.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	deadLetterTopic := "deadletter"

	// Streaming subscription for topic "sendorder" on pubsub component
	// "messages". The given subscription handler is called when a message is
	// received. The  returned `stop` function is used to stop the subscription
	// and close the connection.
	stop, err := client.SubscribeWithHandler(context.Background(),
		daprd.SubscriptionOptions{
			PubsubName:      "pubsub",
			Topic:           "orders",
			DeadLetterTopic: &deadLetterTopic,
		},
		eventHandler,
	)
	if err != nil {
		log.Fatal(err)
	}

	// Another method of streaming subscriptions, this time for the topic "neworder".
	// The returned `sub` object is used to receive messages.
	// `sub` must be closed once it's no longer needed.

	sub, err := client.Subscribe(context.Background(), daprd.SubscriptionOptions{
		PubsubName:      "pubsub",
		Topic:           "orders",
		DeadLetterTopic: &deadLetterTopic,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(">>Created subscription\n")

	for i := 0; i < 3; i++ {
		msg, err := sub.Receive()
		if err != nil {
			log.Fatalf("error receiving message: %v", err)
		}
		log.Printf(">>Received message\n")
		log.Printf("event - PubsubName: %s, Topic: %s, ID: %s, Data: %s\n", msg.PubsubName, msg.Topic, msg.ID, msg.RawData)

		// Use _MUST_ always signal the result of processing the message, else the
		// message will not be considered as processed and will be redelivered or
		// dead lettered.
		if err := msg.Success(); err != nil {
			log.Fatalf("error sending message success: %v", err)
		}
	}

	time.Sleep(time.Second * 100)

	if err := errors.Join(stop(), sub.Close()); err != nil {
		log.Fatal(err)
	}
}

func eventHandler(e *common.TopicEvent) common.SubscriptionResponseStatus {
	log.Printf(">>Received message\n")
	log.Printf("event - PubsubName: %s, Topic: %s, ID: %s, Data: %s\n", e.PubsubName, e.Topic, e.ID, e.Data)
	return common.SubscriptionResponseStatusSuccess
}
