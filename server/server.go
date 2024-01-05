package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/sangharsh/color-sort/db"
	pb "github.com/sangharsh/color-sort/gen/modelpb"
	"github.com/sangharsh/color-sort/level"
	"github.com/sangharsh/color-sort/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type ColorSortApiServer struct {
	pb.UnimplementedColorSortApiServer
}

func (server *ColorSortApiServer) NewLevel(ctx context.Context, req *pb.NewLevelPlayRequest) (*pb.LevelState, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("unable to read context")
	}
	userId := md["colorsort-userid"][0]
	log.Printf("ok: %v, md: %v", ok, md)
	level := level.Generate(req.GetId())
	levelPlay := model.NewLevelPlay(level)
	db.Set(userId, levelPlay)
	return levelPlay.GetCurrentState(), nil
}

func (server *ColorSortApiServer) Pour(ctx context.Context, req *pb.PourRequest) (*pb.PourResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("unable to read context")
	}
	userId := md["colorsort-userid"][0]
	if userId == "" {
		return nil, errors.New("userId not found in request")
	}
	levelPlay := db.Get(userId)

	if levelPlay == nil {
		return nil, errors.New("levelPlay not found for user")
	}
	// log.Printf("Field: %v\n%v\n%v", req.GetSrc(), req.GetDst(), req.GetLevel())
	// return &pb.PourResponse{}, nil
	return model.Pour(req, levelPlay), nil
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
