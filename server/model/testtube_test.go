package model

import (
	"testing"
)

func TestIsCompleteNotFull(t *testing.T) {
	tt := NewTesttube(4, []Color{Color_RED, Color_RED})
	if tt.IsComplete() {
		t.Fatalf(`Half tt returned as completed. tt: %v`, tt)
	}
}

func TestIsCompleteDiffColors(t *testing.T) {
	tt := NewTesttube(4, []Color{Color_RED, Color_RED, Color_RED, Color_GREEN})
	if tt.IsComplete() {
		t.Fatalf(`Diff colors tt returned as completed. tt: %v`, tt)
	}
}

func TestPopEmpty(t *testing.T) {
	tt := NewTesttube(4, []Color{})
	_, err := tt.Pop()
	if  err == nil {
		t.Fatalf(`Popped empty tt. tt: %v; Error: %v`, tt, err)
	}
}

func TestAddColorValid(t *testing.T) {
	tt := NewTesttube(4, []Color{Color_RED, Color_GREEN})
	color := Color_GREEN
	err := tt.AddColor(color)
	if err != nil {
		t.Fatalf(`Unable to add color %q to %v. Error: %v`, color, tt, err)
	}
}

func TestAddColorNonMatching(t *testing.T) {
	tt := NewTesttube(4, []Color{Color_RED, Color_GREEN})
	color := Color_BLUE
	err := tt.AddColor(color)
	if err == nil {
		t.Fatalf(`Added a non-matching color %q to %v.`, color, tt)
	}
}

func TestAddColorFull(t *testing.T) {
	tt := NewTesttube(4, []Color{Color_RED, Color_GREEN, Color_BLUE, Color_YELLOW})
	color := Color_BLUE
	err := tt.AddColor(color)
	if err == nil {
		t.Fatalf(`Added to a full testtube %v.`, tt)
	}
}
