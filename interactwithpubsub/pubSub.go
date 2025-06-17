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
	fmt.Printf("Published: %s to channel: %s\n", message, channel)
}

func SubscribeToChannel(ctx context.Context, client *redis.Client, channel string) {
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
		fmt.Println("listening ", message.Channel, message.Payload)
	}

}
