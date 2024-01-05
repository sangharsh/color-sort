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

func getLevelPlayFromDb(ctx context.Context) (string, *pb.LevelPlay, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", nil, errors.New("unable to read context")
	}
	userId := md["colorsort-userid"][0]
	if userId == "" {
		return "", nil, errors.New("userId not found in request")
	}
	levelPlay := db.Get(userId)

	if levelPlay == nil {
		return "", nil, errors.New("levelPlay not found for user")
	}
	return userId, levelPlay, nil
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
	_, levelPlay, err := getLevelPlayFromDb(ctx)
	if err != nil {
		return nil, err
	}
	return model.Pour(req, levelPlay), nil
}

func (server *ColorSortApiServer) Reset(ctx context.Context, req *pb.ResetRequest) (*pb.LevelState, error) {
	userId, levelPlay, err := getLevelPlayFromDb(ctx)
	if err != nil {
		return nil, err
	}
	levelId := levelPlay.GetCurrentState().GetId()
	level := level.Generate(levelId)
	levelPlayNew := model.NewLevelPlay(level)
	db.Set(userId, levelPlayNew)
	return levelPlayNew.GetCurrentState(), nil
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
