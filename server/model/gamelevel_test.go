package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPourValid(t *testing.T) {
	level := GameLevel{1, []Testtube{Testtube{4, []string{"red", "green"}}, Testtube{4, []string{"red", "green"}}}}

	_, err := level.Pour(0, 1)
	if err != nil {
		t.Fatalf(`Unable to pour. Level: %v \n Error: %v`, level, err)
	}
	assert.Equal(t, []string{"red", "green", "green"}, level.tubes[1].colors, "Tube colors should match")
}

func TestPourNonMatching(t *testing.T) {
	level := GameLevel{1, []Testtube{Testtube{4, []string{"red", "blue"}}, Testtube{2, []string{"red", "green"}}}}

	_, err := level.Pour(0, 1)
	if err == nil {
		t.Fatalf(`Able to pour. Level: %v`, level)
	}
}

func TestPourDstFull(t *testing.T) {
	level := GameLevel{1, []Testtube{Testtube{4, []string{"red", "green"}}, Testtube{2, []string{"red", "green"}}}}

	_, err := level.Pour(0, 1)
	if err == nil {
		t.Fatalf(`Able to pour. Level: %v`, level)
	}
}

func TestPourSrcEmpty(t *testing.T) {
	level := GameLevel{1, []Testtube{Testtube{4, []string{}}, Testtube{2, []string{"red", "green"}}}}

	_, err := level.Pour(0, 1)
	if err == nil {
		t.Fatalf(`Able to pour. Level: %v`, level)
	}
}

func TestGamePlay(t *testing.T) {
	tt := NewTesttube(4, []string{"red", "green", "red", "green"})
	tt2 := NewTesttube(4, []string{"red", "green", "red", "green"})
	gameLevel := NewGameLevel(1, []Testtube{*tt, *tt2})
	solve := [][]int{{0, 2}, {0, 3}, {0, 2}, {0, 3}, {1, 2}, {1, 3}, {1, 2}, {1, 3}}

	for _, move := range solve {
		_, err := gameLevel.Pour(move[0], move[1])
		if err != nil {
			t.Fatalf(`Unable to pour. Error: %v`, err)
		}
	}

	if !gameLevel.Won() {
		t.Fatalf(`Game not won. Level: %v`, gameLevel)
	}
}
