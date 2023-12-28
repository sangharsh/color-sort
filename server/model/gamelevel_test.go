package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPourValid(t *testing.T) {
	level := GameLevel{
		Level: 1,
		Tubes: []*Testtube{{Size: 4, Colors: []Color{Color_RED, Color_GREEN}}, {Size: 4, Colors: []Color{Color_RED, Color_GREEN}}},
	}

	_, err := level.Pour(0, 1)
	if err != nil {
		t.Fatalf(`Unable to pour. Level: %v \n Error: %v`, &level, err)
	}
	assert.Equal(t, []Color{Color_RED, Color_GREEN, Color_GREEN}, level.GetTubes()[1].GetColors(), "Tube colors should match")
}

func TestPourNonMatching(t *testing.T) {
	level := GameLevel{
		Level: 1,
		Tubes: []*Testtube{{Size: 4, Colors: []Color{Color_RED, Color_BLUE}}, {Size: 4, Colors: []Color{Color_RED, Color_GREEN}}},
	}

	_, err := level.Pour(0, 1)
	if err == nil {
		t.Fatalf(`Able to pour. Level: %v`, &level)
	}
}

func TestPourDstFull(t *testing.T) {
	level := GameLevel{
		Level: 1,
		Tubes: []*Testtube{{Size: 4, Colors: []Color{Color_RED, Color_GREEN}}, {Size: 2, Colors: []Color{Color_RED, Color_GREEN}}},
	}
	_, err := level.Pour(0, 1)
	if err == nil {
		t.Fatalf(`Able to pour. Level: %v`, &level)
	}
}

func TestPourSrcEmpty(t *testing.T) {
	level := GameLevel{
		Level: 1,
		Tubes: []*Testtube{{Size: 4, Colors: []Color{}}, {Size: 4, Colors: []Color{Color_RED, Color_GREEN}}},
	}

	_, err := level.Pour(0, 1)
	if err == nil {
		t.Fatalf(`Able to pour. Level: %v`, &level)
	}
}

func TestGamePlay(t *testing.T) {
	level := NewGameLevel(
		1,
		[]*Testtube{{Size: 4, Colors: []Color{Color_RED, Color_GREEN, Color_RED, Color_GREEN}}, {Size: 4, Colors: []Color{Color_RED, Color_GREEN, Color_RED, Color_GREEN}}},
	)

	if level.Won() {
		t.Fatalf(`Game won. Level: %v`, &level)
	}

	solve := [][]int{{0, 2}, {0, 3}, {0, 2}, {0, 3}, {1, 2}, {1, 3}, {1, 2}, {1, 3}}

	for _, move := range solve {
		_, err := level.Pour(move[0], move[1])
		if err != nil {
			t.Fatalf(`Unable to pour. Error: %v`, err)
		}
	}

	if !level.Won() {
		t.Fatalf(`Game not won. Level: %v`, &level)
	}
}
