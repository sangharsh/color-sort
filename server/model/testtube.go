package model

import (
	"errors"

	pb "github.com/sangharsh/color-sort/gen/modelpb"
)

func NewTesttube(size int, colors []pb.Color) *pb.Testtube {
	return &pb.Testtube{
		Size:   int32(size),
		Colors: colors,
	}
}

func IsEmpty(tt *pb.Testtube) bool {
	return len(tt.GetColors()) == 0
}

func IsFull(tt *pb.Testtube) bool {
	return int32(len(tt.GetColors())) == tt.GetSize()
}

func IsComplete(tt *pb.Testtube) bool {
	if IsEmpty(tt) {
		return true
	}
	if !IsFull(tt) {
		return false
	}
	for _, e := range tt.GetColors() {
		if e != tt.GetColors()[0] {
			return false
		}
	}
	return true
}

func Peek(tt *pb.Testtube) (pb.Color, error) {
	if len(tt.GetColors()) == 0 {
		return pb.Color_BLANK, errors.New("tt is empty")
	}
	return tt.GetColors()[len(tt.GetColors())-1], nil
}

func Pop(tt *pb.Testtube) (pb.Color, error) {
	if len(tt.GetColors()) == 0 {
		return pb.Color_BLANK, errors.New("tt is empty")
	}
	color := tt.GetColors()[len(tt.GetColors())-1]
	tt.Colors = tt.GetColors()[:len(tt.GetColors())-1]
	return color, nil
}

func AddColor(tt *pb.Testtube, color pb.Color) error {
	if len(tt.Colors) > 0 && tt.Colors[len(tt.Colors)-1] != color {
		return errors.New("color not matching")
	}
	if len(tt.Colors) == int(tt.GetSize()) {
		return errors.New("tt is full")
	}
	tt.Colors = append(tt.Colors, color)
	return nil
}
