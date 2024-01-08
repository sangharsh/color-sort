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

func getUserFromDb(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", errors.New("unable to read context")
	}
	userId := md["colorsort-userid"][0]
	if userId == "" {
		return "", errors.New("userId not found in request")
	}
	return userId, nil
}

func getLevelPlayFromDb(userId string) *pb.LevelPlay {
	return db.Get(userId)
}

func getUserAndLevelFromDb(ctx context.Context) (string, *pb.LevelPlay, error) {
	userId, err := getUserFromDb(ctx)
	if err != nil {
		return "", nil, err
	}
	levelPlay := getLevelPlayFromDb(userId)
	return userId, levelPlay, nil
}

func (server *ColorSortApiServer) GetLevel(ctx context.Context, req *pb.GetLevelRequest) (*pb.LevelState, error) {
	log.Printf("Entry: \nreq: %v", req.String())
	userId, levelPlay, err := getUserAndLevelFromDb(ctx)
	if err != nil {
		return nil, err
	}

	if levelPlay != nil {
	return levelPlay.GetCurrentState(), nil
	}

	levelId := int32(1)
	level := level.Generate(levelId)
	levelPlayNew := model.NewLevelPlay(level)
	db.Set(userId, levelPlayNew)
	return levelPlayNew.GetCurrentState(), nil
}

func (server *ColorSortApiServer) Pour(ctx context.Context, req *pb.PourRequest) (*pb.PourResponse, error) {
	log.Printf("Entry: \nreq: %v", req.String())
	_, levelPlay, err := getUserAndLevelFromDb(ctx)
	if err != nil {
		return nil, err
	}
	return model.Pour(req, levelPlay), nil
}

func (server *ColorSortApiServer) Reset(ctx context.Context, req *pb.ResetRequest) (*pb.LevelState, error) {
	log.Printf("Reset: \nreq: %v", req.String())
	userId, levelPlay, err := getUserAndLevelFromDb(ctx)
	if err != nil {
		return nil, err
	}
	levelId := levelPlay.GetCurrentState().GetId()
	level := level.Generate(levelId)
	levelPlayNew := model.NewLevelPlay(level)
	db.Set(userId, levelPlayNew)
	return levelPlayNew.GetCurrentState(), nil
}

func (server *ColorSortApiServer) Undo(ctx context.Context, req *pb.UndoRequest) (*pb.LevelState, error) {
	log.Printf("Entry: \nreq: %v", req.String())
	_, levelPlay, err := getLevelPlayFromDb(ctx)
	if err != nil {
		return nil, err
	}

	return model.Undo(req, levelPlay)
}

func (server *ColorSortApiServer) NextLevel(ctx context.Context, req *pb.NextRequest) (*pb.LevelState, error) {
	log.Printf("NextLevel - Entry: \nreq: %v", req.String())
	userId, levelPlay, err := getUserAndLevelFromDb(ctx)
	if err != nil {
		return nil, err
	}
	levelId := levelPlay.GetCurrentState().GetId() + 1
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
