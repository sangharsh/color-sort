package model

import (
	pb "github.com/sangharsh/color-sort/gen/modelpb"
)

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
