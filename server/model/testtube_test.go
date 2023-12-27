package model

import (
	"testing"
)

func TestAddColorValid(t *testing.T) {
	tt := &Testtube{Size: 4, Colors: []Color{Color_RED, Color_GREEN}}
	color := Color_GREEN
	err := tt.AddColor(color)
	if err != nil {
		t.Fatalf(`Unable to add color %q to %v. Error: %v`, color, tt, err)
	}
}

func TestAddColorNonMatching(t *testing.T) {
	tt := &Testtube{Size: 4, Colors: []Color{Color_RED, Color_GREEN}}
	color := Color_BLUE
	err := tt.AddColor(color)
	if err == nil {
		t.Fatalf(`Added a non-matching color %q to %v.`, color, tt)
	}
}

func TestAddColorFull(t *testing.T) {
	tt := &Testtube{Size: 4, Colors: []Color{Color_RED, Color_GREEN, Color_BLUE, Color_YELLOW}}
	color := Color_BLUE
	err := tt.AddColor(color)
	if err == nil {
		t.Fatalf(`Added to a full testtube %v.`, tt)
	}
}
