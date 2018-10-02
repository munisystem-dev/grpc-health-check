package main

import (
	"context"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
)

const (
	address     = "localhost:8080"
	serviceName = "sample"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("didn't connect: %v", err)
	}
	ctx := context.Background()
	defer conn.Close()
	resp, err := grpc_health_v1.NewHealthClient(conn).Check(ctx, &grpc_health_v1.HealthCheckRequest{
		Service: "sample",
	})
	if err != nil {
		if stat, ok := status.FromError(err); ok && stat.Code() == codes.Unimplemented {
			log.Println("the server doesn't implement the grpc health protocol")
		} else {
			log.Printf("rpc failed %s", err)
		}
		os.Exit(1)
	}
	log.Printf("status: %s", resp.GetStatus().String())
}
