package interactwithpubsub

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func PublishMessage(ctx context.Context, client *redis.Client, channel string, message string) {
	err := client.Publish(ctx, channel, message).Err()
	if err != nil {
		panic(err)
	}
	fmt.Println("Published:", message, "to channel:", channel)

}

func SubscribeToChannel(ctx context.Context, client *redis.Client, channel string, msgChan chan string) {
	pubsub := client.Subscribe(ctx, channel)

	_, err := pubsub.Receive(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Subscribed to channel: %s\n", channel)

	for {
		message, err := pubsub.ReceiveMessage(ctx)
		if err != nil {
			panic(err)
		}
		msg := fmt.Sprintf("listening %s: %s", message.Channel, message.Payload)
		msgChan <- msg
	}

}
