package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "ad/ch04/grpcr/proto"
)

func main() {
	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, &server{})

	// Register reflection service on gRPC server.
	reflection.Register(s)

	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}

	s.Serve(lis)
}
