package model

import (
	"errors"
	//"github.com/golang-collections/collections/stack"
)

type Testtube struct {
	size int
	// TODO: Use stack data structure
	colors []string
}

func NewTesttube(size int, colors []string) *Testtube {
	return &Testtube{size, colors}
}

func (tt Testtube) IsEmpty() bool {
	return len(tt.colors) == 0
}

func (tt Testtube) IsFull() bool {
	return len(tt.colors) == tt.size
}

func (tt Testtube) IsComplete() bool {
	if tt.IsEmpty() {
		return true
	}
	if !tt.IsFull() {
		return false
	}
	for _, e := range tt.colors {
		if e != tt.colors[0] {
			return false
		}
	}
	return true
}

func (tt Testtube) Peek() (string, error) {
	if len(tt.colors) == 0 {
		return "", errors.New("tt is empty")
	}
	return tt.colors[len(tt.colors)-1], nil
}

func (tt *Testtube) Pop() (string, error) {
	if len(tt.colors) == 0 {
		return "", errors.New("tt is empty")
	}
	color := tt.colors[len(tt.colors)-1]
	tt.colors = tt.colors[:len(tt.colors)-1]
	return color, nil
}

func (tt *Testtube) AddColor(color string) error {
	if len(tt.colors) == tt.size {
		return errors.New("tt is full")
	}
	if tt.colors[len(tt.colors)-1] != color {
		return errors.New("color not matching")
	}
	tt.colors = append(tt.colors, color)
	return nil
}
