package main

import (
	"context"
	"io"

	pb "ad/ch04/grpcr/proto"
)

type server struct {
	pb.UnimplementedHelloServiceServer
}

func (s *server) Hello(ctx context.Context, in *pb.String) (*pb.String, error) {
	return nil, nil
}

func (s *server) Channel(stream pb.HelloService_ChannelServer) error {
	for {
		args, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		reply := &pb.String{Value: "hello:" + args.GetValue()}

		err = stream.Send(reply)
		if err != nil {
			return err
		}
	}
}
