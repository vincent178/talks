package main

import (
	"context"
	"log"

	v1 "github.com/vincent178/talks/proto-compatibility/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
    conn, err := grpc.NewClient("127.0.0.1:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
    defer conn.Close()

    if err != nil {
        log.Fatalf("failed to dial: %v", err)
    }

    client := v1.NewHelloServiceClient(conn)
    resp, err := client.Hello(context.Background(), &v1.HelloRequest{Name: "Vincent"})
    if err != nil {
        log.Fatalf("failed to call grpc: %v", err)
    }
    log.Printf("resp: %v", resp)
}
