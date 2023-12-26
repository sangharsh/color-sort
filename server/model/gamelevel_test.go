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
