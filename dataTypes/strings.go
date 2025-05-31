package dataTypes

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func SetKeyNX(ctx context.Context, client *redis.Client, key, value string, ttl time.Duration) {
	ok, err := client.SetNX(ctx, key, value, ttl).Result()
	if err != nil {
		panic(err)
	}
	if ok {
		fmt.Println("Key set successfully")
	} else {
		fmt.Println("Key already exists, not set")
	}
}

func GetKey(ctx context.Context, client *redis.Client, key string) {
	val, err := client.Get(ctx, key).Result()
	if err == redis.Nil {
		fmt.Println("Key does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Printf("Key: %s, Value: %s\n", key, val)
	}

}
