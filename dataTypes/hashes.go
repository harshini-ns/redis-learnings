package dataTypes

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func HashSet(ctx context.Context, client *redis.Client, fields []string) error {
	if len(fields)%2 != 0 {
		return fmt.Errorf("invalid number of fields: must be key-value pairs")
	}
	args := make([]interface{}, len(fields))
	for i, field := range fields {
		args[i] = field
	}
	return client.HSet(ctx, "bike:1000", args...).Err()
}

func GetHash(ctx context.Context, client *redis.Client) (map[string]string, error) {
	result, err := client.HGetAll(ctx, "bike:1000").Result()
	if err != nil {
		return nil, fmt.Errorf("error fetching hash: %w", err)
	}
	if len(result) == 0 {
		return nil, fmt.Errorf("hash at bike:1000 not found or is empty")
	}
	return result, nil
}
