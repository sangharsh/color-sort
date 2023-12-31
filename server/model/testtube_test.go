package model

import (
	"testing"

	pb "github.com/sangharsh/color-sort/gen/modelpb"
)

func TestIsCompleteNotFull(t *testing.T) {
	tt := NewTesttube(2, []pb.Color{pb.Color_RED})
	if isComplete(tt) {
		t.Fatalf(`Half tt returned as completed. tt: %v`, tt)
	}
}

func TestIsCompleteDiffColors(t *testing.T) {
	tt := NewTesttube(2, []pb.Color{pb.Color_RED, pb.Color_GREEN})
	if isComplete(tt) {
		t.Fatalf(`Diff pb.Colors tt returned as completed. tt: %v`, tt)
	}
}

func TestPopEmpty(t *testing.T) {
	tt := NewTesttube(2, []pb.Color{})
	_, err := pop(tt)
	if err == nil {
		t.Fatalf(`Popped empty tt. tt: %v; Error: %v`, tt, err)
	}
}

func TestAddColorValid(t *testing.T) {
	tt := NewTesttube(4, []pb.Color{pb.Color_RED, pb.Color_GREEN})
	color := pb.Color_GREEN
	err := addColor(tt, color)
	if err != nil {
		t.Fatalf(`Unable to add pb.Color %q to %v. Error: %v`, color, tt, err)
	}
}

func TestAddColorNonMatching(t *testing.T) {
	tt := NewTesttube(4, []pb.Color{pb.Color_RED, pb.Color_GREEN})
	color := pb.Color_BLUE
	err := addColor(tt, color)
	if err == nil {
		t.Fatalf(`Added a non-matching pb.Color %q to %v.`, color, tt)
	}
}

func TestAddColorFull(t *testing.T) {
	tt := NewTesttube(2, []pb.Color{pb.Color_RED, pb.Color_BLUE})
	color := pb.Color_BLUE
	err := addColor(tt, color)
	if err == nil {
		t.Fatalf(`Added to a full testtube %v.`, tt)
	}
}

func TestForceAddMatchingColor(t *testing.T) {
	tt := NewTesttube(2, []pb.Color{pb.Color_RED})
	color := pb.Color_RED
	err := forceAddColor(tt, color)
	if err != nil {
		t.Fatalf(`Unable to add pb.Color %q to %v. Error: %v`, color, tt, err)
	}
}

func TestForceAddNonMatchingColor(t *testing.T) {
	tt := NewTesttube(4, []pb.Color{pb.Color_RED})
	color := pb.Color_BLUE
	err := forceAddColor(tt, color)
	if err != nil {
		t.Fatalf(`Forcing non-matching color failed. Color %q to %v.`, color, tt)
	}
}

func TestForceAddColorFull(t *testing.T) {
	tt := NewTesttube(2, []pb.Color{pb.Color_RED, pb.Color_GREEN})
	color := pb.Color_BLUE
	err := forceAddColor(tt, color)
	if err == nil {
		t.Fatalf(`Added to a full testtube %v.`, tt)
	}
}
