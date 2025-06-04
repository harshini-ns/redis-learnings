package main

import (
	"context"
	"fmt"
	"log"
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
	dataTypes.SetKeyXX(ctx, client, "fop", "jaan", 10*time.Second)
	dataTypes.GetKey(ctx, client, "fop")

	//push is done
	listKey := "bikes:repairs"

	dataTypes.PushToList(ctx, client, listKey, "bike:19", "bike:29", "bike:38")
	//get elements from the push action
	elements, err := dataTypes.GetListElements(ctx, client, listKey)
	if err != nil {
		log.Fatalf("Error retrieving list elements: %v", err)
	}
	fmt.Println("List elements:", elements)

	//pop is done
	val, err := dataTypes.PopFromList(ctx, client, listKey)
	if err != nil {
		log.Fatalf("Error popping from list: %v", err)
	}
	fmt.Println("Popped value:", val)
	//get elements after pop action
	ele, err := dataTypes.GetListElements(ctx, client, listKey)
	if err != nil {
		log.Fatalf("Error retrieving list elements: %v", err)
	}
	fmt.Println("after poppping:", ele)

	go PushNumberstoList(ctx, client, listKey)
	//BRPOP
	for i := 0; i < 5; i++ {
		fmt.Println("BRPOP:Waiting for element...")
		value, err := dataTypes.BlockRpop(ctx, client, listKey)
		if err != nil {
			log.Fatalf("Error during BRPOP: %v", err)
		}

		if value == "" {
			fmt.Println("BRPOP :no element retrieved, timed out.")
		} else {
			fmt.Printf("brpop: Popped value: %s\n", value)
		}
	}

}

func PushNumberstoList(ctx context.Context, client *redis.Client, listKey string) {
	for i := 0; i <= 5; i++ {
		fmt.Println("Pushing bikes to list...")
		dataTypes.PushToList(ctx, client, listKey, "bike:10", "bike:88", "bike:90")
		time.Sleep(20 * time.Second)
	}
}
