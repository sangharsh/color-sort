package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/sangharsh/color-sort/model"
	pb "github.com/sangharsh/color-sort/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type ColorSortApiServer struct {
	pb.UnimplementedColorSortApiServer
}

func (server *ColorSortApiServer) GetGameLevel(ctx context.Context, req *model.LevelRequest) (*model.GameLevel, error) {
	level := model.NewGameLevel(
		1,
		[]*model.Testtube{{Size: 4, Colors: []model.Color{model.Color_RED, model.Color_GREEN, model.Color_RED, model.Color_GREEN}}, {Size: 4, Colors: []model.Color{model.Color_RED, model.Color_GREEN, model.Color_RED, model.Color_GREEN}}},
	)
	return level, nil
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
