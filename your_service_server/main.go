package main

import (
	"context"
	"google.golang.org/grpc"
	pb "grpcGatewayTest/gen/go"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedYourServiceServer
}

func (s *server) Echo(ctx context.Context, in *pb.StringMessage) (*pb.StringMessage, error) {
	log.Printf("Received: %v", in.GetValue())
	//return in, nil
	return &pb.StringMessage{Value: "_" + in.GetValue()}, nil
}
func main() {
	lis, err := net.Listen("tcp", ":50001")
	if err != nil {
		log.Fatalf("failed to listen : %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterYourServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve : %v", err)
	}

}
