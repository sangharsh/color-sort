package model

import (
	pb "github.com/sangharsh/color-sort/gen/modelpb"
)

func NewGameLevel(level int32, tubes []*pb.Testtube) *pb.GameLevel {
	tubes2 := append(tubes,
		NewTesttube(4, []pb.Color{}),
		NewTesttube(4, []pb.Color{}))
	return &pb.GameLevel{
		Level: level,
		Tubes: tubes2,
	}
}

func pour(level *pb.GameLevel, srcidx int, dstidx int) (bool, error) {
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

func won(level *pb.GameLevel) bool {
	for _, tt := range level.GetTubes() {
		if !isComplete(tt) {
			return false
		}
	}
	return true
}
