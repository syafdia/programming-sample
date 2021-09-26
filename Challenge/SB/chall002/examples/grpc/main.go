package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"

	grpcdelivery "github.com/syafdia/sb/chal002/internal/delivery/grpc"
)

func main() {
	address := "0.0.0.0:8002"
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("can not connect to server, err: %v", err)
	}

	defer conn.Close()

	movieClient := grpcdelivery.NewMovieServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Find multiple movie.
	response, err := movieClient.FindMultiple(ctx, &grpcdelivery.MovieFindMultipleRequest{
		SearchWord: "Batman",
		Pagination: 1,
	})
	if err != nil {
		log.Fatalf("failed on FindMultiple, err: %v", err)
	}

	log.Println("got response from FindMultiple, response:", response)
}
