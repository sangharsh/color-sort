package model

import (
	"errors"

	pb "github.com/sangharsh/color-sort/gen/modelpb"
)

func NewLevel(levelId int32, tubes []*pb.Testtube) *pb.LevelState {
	tubes2 := append(tubes,
		NewTesttube(4, []pb.Color{}),
		NewTesttube(4, []pb.Color{}))
	level := &pb.LevelState{
		Id:    levelId,
		Tubes: tubes2,
	}
	level.Won = won(level)
	return level
}

type addFn func(*pb.Testtube, pb.Color) error

// TODO: Moves only a single item right now
func move(from *pb.Testtube, to *pb.Testtube, fn addFn) (bool, error) {
	color, err := peek(from)
	if err != nil {
		return false, err
	}
	err = fn(to, color)
	if err != nil {
		return false, err
	}
	_, err = pop(from)
	if err != nil {
		return false, err
	}

	return true, nil
}

// TODO: Pours only a single item right now
func pour(level *pb.LevelState, srcidx int, dstidx int) (bool, error) {
	if level.Won {
		return false, errors.New("level has been won")
	}
	src := level.Tubes[srcidx]
	dst := level.Tubes[dstidx]
	ok, err := move(src, dst, addColor)
	if !ok || err != nil {
		return ok, err
	}
	level.Won = won(level)
	return true, nil
}

func undo(level *pb.LevelState, moveResp *pb.PourSuccessResponse) (bool, error) {
	if level.Won {
		return false, errors.New("level has been won")
	}
	src := level.Tubes[moveResp.GetSrc()]
	dst := level.Tubes[moveResp.GetDst()]
	// TODO: numItemsPoured := move.GetNumItemsPoured()
	return move(dst, src, forceAddColor)
}

func won(level *pb.LevelState) bool {
	for _, tt := range level.GetTubes() {
		if !isComplete(tt) {
			return false
		}
	}
	return true
}
