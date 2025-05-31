package main

import (
	"context"
	"time"

	"todo-app-redis-go/dataTypes"

	"github.com/redis/go-redis/v9"
)

func connectRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // No password set
		DB:       0,  // Use default DB
		Protocol: 2,  // Connection protocol
	})
	return client
}

func main() {
	client := connectRedis()
	ctx := context.Background()
	dataTypes.SetKeyNX(ctx, client, "fop", "bar", 10*time.Second)
	dataTypes.GetKey(ctx, client, "fwop")
}
