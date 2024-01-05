package model

import (
	"fmt"
	"testing"

	pb "github.com/sangharsh/color-sort/gen/modelpb"
	"github.com/stretchr/testify/assert"
)

func TestPourSuccess(t *testing.T) {
	level := NewLevel(1, []*pb.Testtube{
		NewTesttube(4, []pb.Color{pb.Color_RED, pb.Color_GREEN, pb.Color_RED, pb.Color_GREEN}),
		NewTesttube(4, []pb.Color{pb.Color_RED, pb.Color_GREEN, pb.Color_RED, pb.Color_GREEN}),
	})
	levelPlay := NewLevelPlay(level)
	req := &pb.PourRequest{
		Src: 0,
		Dst: 2,
	}
	resp := Pour(req, levelPlay)
	fmt.Printf("PourResponse: %v\n", resp)
	if resp.GetErr() != "" || resp.GetResponse() == nil {
		t.Fatalf("Got error or empty response.\nPourResponse: %v", resp)
	}
	sResp := resp.GetResponse()
	if sResp.Src != req.Src || sResp.Dst != req.Dst || sResp.NumItemsPoured != 1 {
		t.Fatalf("Unexpected src/dst.\nPourResponse: %v", resp)
	}
	move := levelPlay.GetMoves()[0]
	if move.Src != req.Src || move.Dst != req.Dst || move.NumItemsPoured != 1 {
		t.Fatalf("Unexpected move.\nPourResponse: %v", resp)
	}
	assert.Equal(t,
		[]pb.Color{pb.Color_RED, pb.Color_GREEN, pb.Color_RED},
		levelPlay.CurrentState.GetTubes()[0].GetColors(),
		"Src colors should match")
	assert.Equal(t,
		[]pb.Color{pb.Color_GREEN},
		levelPlay.CurrentState.GetTubes()[2].GetColors(),
		"Dst colors should match")
}

func TestPourErr(t *testing.T) {
	level := NewLevel(1, []*pb.Testtube{
		NewTesttube(4, []pb.Color{pb.Color_RED, pb.Color_GREEN, pb.Color_RED, pb.Color_GREEN}),
		NewTesttube(4, []pb.Color{pb.Color_RED, pb.Color_GREEN, pb.Color_RED, pb.Color_GREEN}),
	})
	levelPlay := NewLevelPlay(level)
	req := &pb.PourRequest{
		Src: 0,
		Dst: 1,
	}
	resp := Pour(req, levelPlay)
	if resp.GetErr() == "" || resp.GetResponse() != nil {
		t.Fatalf("Got empty error or valid response.\nPourResponse: %v", resp)
	}
}
func TestGamePlay(t *testing.T) {
	level := NewLevel(1, []*pb.Testtube{
		NewTesttube(4, []pb.Color{pb.Color_RED, pb.Color_GREEN, pb.Color_RED, pb.Color_GREEN}),
		NewTesttube(4, []pb.Color{pb.Color_RED, pb.Color_GREEN, pb.Color_RED, pb.Color_GREEN}),
	})
	levelPlay := NewLevelPlay(level)

	if won(levelPlay.GetCurrentState()) {
		t.Fatalf(`Game won. Level: %v`, &level)
	}

	solve := [][]int32{{0, 2}, {0, 3}, {0, 2}, {0, 3}, {1, 2}, {1, 3}, {1, 2}, {1, 3}}

	for _, move := range solve {
		pourResp := Pour(&pb.PourRequest{
			Src: move[0],
			Dst: move[1],
		}, levelPlay)
		// _, err := pour(level, move[0], move[1])
		if pourResp.GetErr() != "" {
			t.Fatalf(`Unable to pour. PourResponse: %v`, pourResp)
		}
	}

	if !won(levelPlay.GetCurrentState()) {
		t.Fatalf(`Game not won. Level: %v`, &level)
	}
}

func TestUndoSuccess(t *testing.T) {
	tubes := []*pb.Testtube{
		NewTesttube(2, []pb.Color{pb.Color_RED}),
		NewTesttube(2, []pb.Color{pb.Color_GREEN}),
	}
	levelPlay := &pb.LevelPlay{
		CurrentState: &pb.LevelState{
			Id:    1,
			Tubes: tubes,
		},
		Moves: []*pb.PourSuccessResponse{
			{Src: 0, Dst: 1, NumItemsPoured: 1},
		},
	}
	req := &pb.UndoRequest{}
	level, err := Undo(req, levelPlay)

	if err != nil {
		t.Fatalf("Got error.\nErr: %v", err)
	}
	assert.Equal(t,
		[]pb.Color{pb.Color_RED, pb.Color_GREEN},
		level.GetTubes()[0].GetColors(),
		"src colors should match")
	assert.Equal(t,
		[]pb.Color{},
		levelPlay.CurrentState.GetTubes()[1].GetColors(),
		"dst colors should match")
}
