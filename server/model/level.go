package model

import (
	pb "github.com/sangharsh/color-sort/gen/modelpb"
)

func NewLevel(level int32, tubes []*pb.Testtube) *pb.Level {
	tubes2 := append(tubes,
		NewTesttube(4, []pb.Color{}),
		NewTesttube(4, []pb.Color{}))
	return &pb.Level{
		Id:         level,
		StartState: &pb.LevelState{Tubes: tubes2},
	}
}
