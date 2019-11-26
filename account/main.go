package main

import (
	"context"
	"log"
	"time"

	"github.com/sauravgsh16/go-store/client"
)

func main() {
	c, err := client.NewClient("localhost:8080")
	if err != nil {
		log.Fatalf("Fail to connect to client")
	}
	defer c.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	acc, err := c.PostAccount(ctx, "First test")
	if err != nil {
		log.Fatalf("Failed to post: %s", err.Error())
	}
	log.Printf("Post result: %+v\n", acc)

	acc, err = c.GetAccount(ctx, acc.Id)
	if err != nil {
		log.Fatalf("failed to get: %s", err.Error())
	}
	log.Printf("Get result: %+v\n", acc)

	accs, err := c.GetAccounts(ctx, uint64(0), uint64(0))
	if err != nil {
		log.Fatalf("failed to get accounts: %s", err.Error())
	}
	log.Printf("Get accounts result: %+v\n", accs)
}
