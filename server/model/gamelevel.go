package model

import (
	pb "github.com/sangharsh/color-sort/gen/modelpb"
)

type GameLevel struct {
	Glpb pb.GameLevel
}

func NewGameLevel(level int32, tubes []*pb.Testtube) *GameLevel {
	tubes2 := append(tubes, NewTesttube(4, []pb.Color{}), NewTesttube(4, []pb.Color{}))
	gameLevel := &GameLevel{pb.GameLevel{
		Level: level,
		Tubes: tubes2,
	}}
	return gameLevel
}

func (level *GameLevel) Pour(srcidx, dstidx int) (bool, error) {
	src := level.Glpb.Tubes[srcidx]
	dst := level.Glpb.Tubes[dstidx]
	color, err := Peek(src)
	if err != nil {
		return false, err
	}
	err = AddColor(dst, color)
	if err != nil {
		return false, err
	}
	_, err = Pop(src)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (level *GameLevel) Won() bool {
	for _, tt := range level.Glpb.GetTubes() {
		if !IsComplete(tt) {
			return false
		}
	}
	return true
}
