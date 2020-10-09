// The client does an end-to-end test with the server.
package main

import (
	"context"
	pb "github.com/j0holo/keyver/rpc"
	"google.golang.org/grpc"
	"log"
)

const port = ":8181"

func main() {
	conn, err := grpc.Dial("localhost" + port, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := pb.NewDBClient(conn)
	ctx := context.Background()

	r, err := c.Set(ctx, &pb.SetRequest{
		Key:   "hello",
		Value: "world",
	})
	if err != nil {
		log.Println(err)
	}

	log.Println(r.Result)

	r, err = c.Get(ctx, &pb.GetRequest{
		Key:   "hello",
	})
	if err != nil {
		log.Println(err)
	}

	log.Println(r.Result)

	r, err = c.Get(ctx, &pb.GetRequest{
		// This is not allowed, too meta for this code.
		Key:   "key",
	})
	if err != nil {
		log.Println(err)
	} else {
		log.Println(r.Result)
	}
}