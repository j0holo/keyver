package main

import (
	"context"
	pb "github.com/j0holo/keyver/rpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
)

const port = ":8181"

type server struct {
	pb.UnimplementedDBServer
}

func (s *server) Get(ctx context.Context, in *pb.GetRequest) (*pb.Response, error) {
	return handleGet(ctx, in.Key)
}

func (s *server) Set(ctx context.Context, in *pb.SetRequest) (*pb.Response, error) {
	return handleSet(ctx, in.Key, in.Value)
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterDBServer(s, &server{})
	log.Println("Listening on port", port)
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

func handleGet(ctx context.Context, key string) (*pb.Response, error) {
	if key == "key" {
		return nil, status.Errorf(codes.Unavailable, "We don't know the key named '%s'", key)
	}
	log.Printf("Got the set request for %s", key)
	return &pb.Response{Result: "value"}, nil
}

func handleSet(ctx context.Context, key, value string) (*pb.Response, error) {
	log.Printf("Got the get request for %s: %s", key, value)
	return &pb.Response{Result: "value"}, nil
}
