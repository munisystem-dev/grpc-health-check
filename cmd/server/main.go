package main

import (
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

const (
	serviceName = "example"
	port        = ":8080"
)

var isHealth = false

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	server := health.NewServer()
	grpc_health_v1.RegisterHealthServer(s, server)
	go func() {
		for {
			if isHealth {
				server.SetServingStatus(serviceName, grpc_health_v1.HealthCheckResponse_NOT_SERVING)
				isHealth = false
				log.Printf("service %s isn't serving", serviceName)
			} else {
				server.SetServingStatus(serviceName, grpc_health_v1.HealthCheckResponse_SERVING)
				isHealth = true
				log.Printf("service %s is serving", serviceName)
			}
			time.Sleep(time.Second * 5)
		}
	}()
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
