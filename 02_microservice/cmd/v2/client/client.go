package main

import (
	"context"
	"fmt"
	"time"

	pb "github.com/henrjan/microservice/internal/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "localhost:8080"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		fmt.Printf("did not connect: %v\n", err)
	}
	defer conn.Close()
	c := pb.NewSearchMovieClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	req := &pb.SearchRequest{
		SearchWord: "Batman",
		Page:       "1",
	}

	r, err := c.GetMovie(ctx, req)
	if err != nil {
		fmt.Printf("could not greet: %v\n", err)
	}

	fmt.Println(r)
}
