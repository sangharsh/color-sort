package model

import (
	"testing"
)

func TestAddColorValid(t *testing.T) {
	tt := Testtube{4, []string{"red", "green"}}
	color := "green"
	err := tt.AddColor(color)
	if err != nil {
		t.Fatalf(`Unable to add color %q to %v. Error: %v`, color, tt, err)
	}
}

func TestAddColorNonMatching(t *testing.T) {
	tt := Testtube{4, []string{"red", "green"}}
	color := "blue"
	err := tt.AddColor(color)
	if err == nil {
		t.Fatalf(`Added a non-matching color %q to %v.`, color, tt)
	}
}

func TestAddColorFull(t *testing.T) {
	tt := Testtube{4, []string{"red", "green", "blue", "yellow"}}
	color := "blue"
	err := tt.AddColor(color)
	if err == nil {
		t.Fatalf(`Added to a full testtube %v.`, tt)
	}
}
