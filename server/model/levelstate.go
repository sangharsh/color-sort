package model

import (
	pb "github.com/sangharsh/color-sort/gen/modelpb"
)

func NewLevel(level int32, tubes []*pb.Testtube) *pb.LevelState {
	tubes2 := append(tubes,
		NewTesttube(4, []pb.Color{}),
		NewTesttube(4, []pb.Color{}))
	return &pb.LevelState{
		Id:    level,
		Tubes: tubes2,
	}
}

// TODO: Pours only a single item right now
func pour(level *pb.LevelState, srcidx int, dstidx int) (bool, error) {
	src := level.Tubes[srcidx]
	dst := level.Tubes[dstidx]
	color, err := peek(src)
	if err != nil {
		return false, err
	}
	err = addColor(dst, color)
	if err != nil {
		return false, err
	}
	_, err = pop(src)
	if err != nil {
		return false, err
	}

	return true, nil
}

func won(level *pb.LevelState) bool {
	for _, tt := range level.GetTubes() {
		if !isComplete(tt) {
			return false
		}
	}
	return true
}
