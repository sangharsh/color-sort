package model

import (
	"errors"
)

func NewTesttube(size int, colors []Color) *Testtube {
	return &Testtube{
		Size:   int32(size),
		Colors: colors,
	}
}

func (tt *Testtube) IsEmpty() bool {
	return len(tt.GetColors()) == 0
}

func (tt *Testtube) IsFull() bool {
	return int32(len(tt.GetColors())) == tt.GetSize()
}

func (tt *Testtube) IsComplete() bool {
	if tt.IsEmpty() {
		return true
	}
	if !tt.IsFull() {
		return false
	}
	for _, e := range tt.GetColors() {
		if e != tt.GetColors()[0] {
			return false
		}
	}
	return true
}

func (tt *Testtube) Peek() (Color, error) {
	if len(tt.GetColors()) == 0 {
		return Color_BLANK, errors.New("tt is empty")
	}
	return tt.GetColors()[len(tt.GetColors())-1], nil
}

func (tt *Testtube) Pop() (Color, error) {
	if len(tt.GetColors()) == 0 {
		return Color_BLANK, errors.New("tt is empty")
	}
	color := tt.GetColors()[len(tt.GetColors())-1]
	tt.Colors = tt.GetColors()[:len(tt.GetColors())-1]
	return color, nil
}

func (tt *Testtube) AddColor(color Color) error {
	if len(tt.Colors) > 0 && tt.Colors[len(tt.Colors)-1] != color {
		return errors.New("color not matching")
	}
	if len(tt.Colors) == int(tt.GetSize()) {
		return errors.New("tt is full")
	}
	tt.Colors = append(tt.Colors, color)
	return nil
}
