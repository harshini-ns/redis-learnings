package dataTypes

import (
	"context"
	"fmt"
	"time"

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

func GetListElements(ctx context.Context, client *redis.Client, key string) ([]string, error) {
	elements, err := client.LRange(ctx, key, 0, -1).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve list elements: %w", err)
	}
	return elements, nil
}

func PopFromList(ctx context.Context, client *redis.Client, key string) (string, error) {
	val, err := client.LPop(ctx, key).Result()
	if err == redis.Nil {
		return "", nil
	} else if err != nil {
		return "", fmt.Errorf("failed to pop from list: %w", err)
	}
	return val, nil
}

func BlockRpop(ctx context.Context, client *redis.Client, key string) (string, error) {
	timeout := 10 * time.Second
	result, err := client.BRPop(ctx, timeout, key).Result()

	if err != nil {
		if err == redis.Nil {
			// Timeout occurred with no new elements
			return "", nil
		}
		return "", fmt.Errorf("failed to BRPOP from list '%s': %w", key, err)
	}
	return result[1], nil
}
