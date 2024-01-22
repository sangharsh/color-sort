package solver

import (
	"testing"

	pb "github.com/sangharsh/color-sort/gen/modelpb"
	"github.com/sangharsh/color-sort/level"
	"github.com/sangharsh/color-sort/model"
)

func TestRandomSolveSuccess(t *testing.T) {
	level := model.NewLevel(1, []*pb.Testtube{
		model.NewTesttube(4, []pb.Color{pb.Color_RED, pb.Color_GREEN, pb.Color_RED, pb.Color_GREEN}),
		model.NewTesttube(4, []pb.Color{pb.Color_RED, pb.Color_GREEN, pb.Color_RED, pb.Color_GREEN}),
	})
	levelPlay := model.NewLevelPlay(level)

	RandomSolve(levelPlay)
	if !levelPlay.GetCurrentState().GetWon() {
		t.Logf(`Game not won. Level: %v`, levelPlay.GetCurrentState().String())
	}
}

func TestBFSolver(t *testing.T) {
	level := model.NewLevel(1, []*pb.Testtube{
		model.NewTesttube(4, []pb.Color{pb.Color_RED, pb.Color_GREEN, pb.Color_RED, pb.Color_GREEN}),
		model.NewTesttube(4, []pb.Color{pb.Color_RED, pb.Color_GREEN, pb.Color_RED, pb.Color_GREEN}),
	})

	solver := NewBFSSolver(level)
	levelPlay := solver.Solve()

	if !levelPlay.GetCurrentState().GetWon() {
		t.Fatalf(`Game not won. Level: %v`, levelPlay.GetCurrentState().String())
	} else {
		t.Logf("LevelPlay: %v", levelPlay)
		t.Logf("Moves: %v", MovesString(levelPlay))
	}
}

func TestDFSolver(t *testing.T) {

	level := level.Generate(2)

	solver := NewDFSSolver(level)
	levelPlay := solver.Solve()

	if !levelPlay.GetCurrentState().GetWon() {
		t.Fatalf(`Game not won. Level: %v`, levelPlay.GetCurrentState().String())
	} else {
		t.Logf("LevelPlay: %v", levelPlay)
		t.Logf("Moves: %v", MovesString(levelPlay))
	}
}
