package model

import (
	"errors"
	"fmt"

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

func moveOne(from *pb.Testtube, to *pb.Testtube, fn addFn) (bool, error) {
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

func move(from *pb.Testtube, to *pb.Testtube, maxItems int, fn addFn) (int, error) {
	ok, err := moveOne(from, to, fn)
	if !ok || err != nil {
		return 0, err
	}
	numItemsPoured := 1
	for ok && numItemsPoured < maxItems {
		ok, _ = moveOne(from, to, fn)
		if ok {
			numItemsPoured += 1
		}
	}
	return numItemsPoured, nil
}

func pour(level *pb.LevelState, srcidx int, dstidx int) (int, error) {
	if level.Won {
		return 0, errors.New("level has been won")
	}
	src := level.Tubes[srcidx]
	dst := level.Tubes[dstidx]
	numItemsPoured, err := move(src, dst, len(src.GetColors()), addColor)
	if err != nil {
		return numItemsPoured, err
	}
	level.Won = won(level)
	return numItemsPoured, nil
}

func undo(level *pb.LevelState, moveResp *pb.PourSuccessResponse) (int, error) {
	if level.Won {
		return 0, errors.New("level has been won")
	}
	src := level.Tubes[moveResp.GetSrc()]
	dst := level.Tubes[moveResp.GetDst()]
	numItemsPoured := moveResp.GetNumItemsPoured()
	return move(dst, src, int(numItemsPoured), forceAddColor)
}

func won(level *pb.LevelState) bool {
	for _, tt := range level.GetTubes() {
		if !isComplete(tt) {
			return false
		}
	}
	return true
}

func HasValidColorAndTubes(level *pb.LevelState) (bool, error) {
	// Check for id and won field as well
	tubes := level.GetTubes()
	numTubes := len(tubes)
	colorMap := make(map[pb.Color]int)

	for _, tt := range tubes {
		for _, color := range tt.GetColors() {
			colorMap[color]++
		}
	}
	if len(colorMap)+2 != numTubes {
		return false, fmt.Errorf("no. of tubes(%v) is not numColor(%v) + 2", numTubes, len(colorMap))
	}

	for k, v := range colorMap {
		if v != 4 {
			return false, fmt.Errorf("no. of elements is not 4: %v", k)
		}
	}
	return true, nil
}
