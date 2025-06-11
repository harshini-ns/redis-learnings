package dataTypes

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func SetADD(ctx context.Context, client *redis.Client, key string, values ...string) {
	vals := make([]interface{}, len(values))
	for i, v := range values {
		vals[i] = v
	}
	count, err := client.SAdd(ctx, key, vals...).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("SADD : %d element(s) added to list '%s'\n", count, key)
}

func GetListElementsFromSetAdd(ctx context.Context, client *redis.Client, key string) ([]string, error) {
	elements, err := client.SMembers(ctx, key).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve list elements: %w", err)
	}
	return elements, nil
}
