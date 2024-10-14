package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/vincent178/talks/proto-compatibility/v5"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	_ "google.golang.org/grpc/status"
)

type helloService struct {
    pb.UnimplementedHelloServiceServer
}

func (s *helloService) Hello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
    // name := req.Name
    return &pb.HelloReply{
        Message: fmt.Sprintf("Hello %s", "Kang"),
        Version: "v6",
    }, nil
}

func main() {
    lis, err := net.Listen("tcp", ":8080")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    s := grpc.NewServer()

    pb.RegisterHelloServiceServer(s, &helloService{})
    reflection.Register(s)

    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
