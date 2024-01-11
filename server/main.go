package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	"github.com/sangharsh/color-sort/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

func getPort() int {
	portStr, ok := os.LookupEnv("PORT")
	if !ok {
		portStr = "50051"
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		port = 50051
	}
	return port
}

func main() {
	port := getPort()
	log.Printf("Starting server at port %v", port)
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	api.Register(grpcServer)
	// Register reflection service on gRPC server.
	reflection.Register(grpcServer)

	healthcheck := health.NewServer()
	healthpb.RegisterHealthServer(grpcServer, healthcheck)
	healthcheck.SetServingStatus("", healthpb.HealthCheckResponse_SERVING)
	grpcServer.Serve(lis)
}
