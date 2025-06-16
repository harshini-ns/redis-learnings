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

	// Wait for confirmation that subscription is created
	_, err := pubsub.Receive(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Subscribed to channel: %s\n", channel)

}
