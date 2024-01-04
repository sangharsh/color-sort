package model

import (
	pb "github.com/sangharsh/color-sort/gen/modelpb"
)

func NewLevelPlay(level *pb.Level) *pb.LevelPlay {
	return &pb.LevelPlay{
		Level:        level,
		Moves:        []*pb.PourSuccessResponse{},
		CurrentState: &pb.LevelState{Tubes: level.GetStartState().GetTubes()},
	}
}
