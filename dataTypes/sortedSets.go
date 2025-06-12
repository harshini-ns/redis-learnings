package dataTypes

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func SortedSetAdd(ctx context.Context, client *redis.Client, key string) {
	count, err := client.ZAdd(ctx, "racer_scores", redis.Z{Member: "usa", Score: 10}, redis.Z{Member: "Prickett", Score: 14},
		redis.Z{Member: "Castilla", Score: 12}).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("ZAdd : %d elements added to sorted set '%s'\n", count, key)
}

func GetSortedElementsFromSortedAdd(ctx context.Context, client *redis.Client, key string) {
	elements, err := client.ZRangeWithScores(ctx, key, 0, -1).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Sorted elements in '%s':\n", key)
	for i := 0; i < len(elements); i++ {
		z := elements[i]
		fmt.Printf("Member: %s, Score: %.0f\n", z.Member, z.Score)
	}

}

func GetReverseFromSortedSets(ctx context.Context, Client *redis.Client, key string) {
	ele, err := Client.ZRevRange(ctx, key, 0, -1).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Sorted elements in REVERSE '%s':\n", key)
	for i := 0; i < len(ele); i++ {
		z := ele[i]
		fmt.Printf("Member: %s\n", z)
	}
}
