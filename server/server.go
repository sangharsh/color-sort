package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/sangharsh/color-sort/gen/modelpb"
	"github.com/sangharsh/color-sort/level"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type ColorSortApiServer struct {
	pb.UnimplementedColorSortApiServer
}

func (server *ColorSortApiServer) GetGameLevel(ctx context.Context, req *pb.LevelRequest) (*pb.GameLevel, error) {
	level := level.Generate(req.GetLevel())
	return &level.Glpb, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterColorSortApiServer(grpcServer, &ColorSortApiServer{})
	// Register reflection service on gRPC server.
	reflection.Register(grpcServer)
	grpcServer.Serve(lis)
}
