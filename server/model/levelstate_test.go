package model

import (
	"testing"

	pb "github.com/sangharsh/color-sort/gen/modelpb"
	"github.com/stretchr/testify/assert"
)

func TestPourValid(t *testing.T) {
	level := &pb.LevelState{
		Tubes: []*pb.Testtube{
			NewTesttube(4, []pb.Color{pb.Color_RED, pb.Color_GREEN}),
			NewTesttube(4, []pb.Color{pb.Color_RED, pb.Color_GREEN}),
		},
	}

	_, err := pour(level, 0, 1)
	if err != nil {
		t.Fatalf(`Unable to pour. Level: %v \n Error: %v`, &level, err)
	}
	assert.Equal(t, []pb.Color{pb.Color_RED, pb.Color_GREEN, pb.Color_GREEN}, level.GetTubes()[1].GetColors(), "Tube colors should match")
}

func TestPourNonMatching(t *testing.T) {
	level := &pb.LevelState{
		Tubes: []*pb.Testtube{
			NewTesttube(4, []pb.Color{pb.Color_RED, pb.Color_BLUE}),
			NewTesttube(4, []pb.Color{pb.Color_RED, pb.Color_GREEN}),
		},
	}

	_, err := pour(level, 0, 1)
	if err == nil {
		t.Fatalf(`Able to pour. Level: %v`, &level)
	}
}

func TestPourDstFull(t *testing.T) {
	level := &pb.LevelState{
		Tubes: []*pb.Testtube{
			NewTesttube(4, []pb.Color{pb.Color_RED, pb.Color_GREEN}),
			NewTesttube(2, []pb.Color{pb.Color_RED, pb.Color_GREEN}),
		},
	}
	_, err := pour(level, 0, 1)
	if err == nil {
		t.Fatalf(`Able to pour. Level: %v`, &level)
	}
}

func TestPourSrcEmpty(t *testing.T) {
	level := &pb.LevelState{
		Tubes: []*pb.Testtube{
			NewTesttube(4, []pb.Color{}),
			NewTesttube(4, []pb.Color{pb.Color_RED, pb.Color_GREEN}),
		},
	}
	_, err := pour(level, 0, 1)
	if err == nil {
		t.Fatalf(`Able to pour. Level: %v`, &level)
	}
}