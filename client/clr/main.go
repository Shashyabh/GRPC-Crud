package main

import (
	"context"
	"log"
	"time"

	"github.com/shashyabh/mygo/client"
)

func main() {
	grpcAddress := "localhost:50051"

	// Create a gRPC client
	c, err := client.NewClient(grpcAddress)
	if err != nil {
		log.Fatalf("Failed to create gRPC client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Call CreateUser
	user, err := c.CreateUser(ctx, "BOB", 50000, "Engineering", "")
	if err != nil {
		log.Fatalf("Failed to create user: %v", err)
	}
	log.Printf("User created: %+v", user)

	// Call GetUser
	// fetchedUser, err := c.GetUser(ctx, user.ID)
	// if err != nil {
	// 	log.Fatalf("Failed to fetch user: %v", err)
	// }
	// log.Printf("Fetched User: %+v", fetchedUser)
}
