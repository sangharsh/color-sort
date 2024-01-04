package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/sangharsh/color-sort/gen/modelpb"
	"github.com/sangharsh/color-sort/level"
	"github.com/sangharsh/color-sort/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type ColorSortApiServer struct {
	pb.UnimplementedColorSortApiServer
}

func (server *ColorSortApiServer) NewLevel(ctx context.Context, req *pb.NewLevelPlayRequest) (*pb.LevelPlay, error) {
	log.Printf("Request: %v\n", req)
	level := level.Generate(req.GetId())
	levelPlay := model.NewLevelPlay(level)
	return levelPlay, nil
}

func (server *ColorSortApiServer) Pour(ctx context.Context, req *pb.PourRequest) (*pb.PourResponse, error) {
	log.Printf("Request: %v\n", req)
	// log.Printf("Field: %v\n%v\n%v", req.GetSrc(), req.GetDst(), req.GetLevel())
	// return &pb.PourResponse{}, nil
	return model.Pour(req), nil
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
