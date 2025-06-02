package dataTypes

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func PushToList(ctx context.Context, client *redis.Client, key string, values ...string) {
	vals := make([]interface{}, len(values))
	for i, v := range values {
		vals[i] = v
	}

	count, err := client.LPush(ctx, key, vals...).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("LPUSH: %d element(s) added to list '%s'\n", count, key)
}
