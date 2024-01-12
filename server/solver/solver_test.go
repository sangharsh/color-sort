package solver

import (
	"testing"

	pb "github.com/sangharsh/color-sort/gen/modelpb"
	"github.com/sangharsh/color-sort/model"
)

func TestSolveSuccess(t *testing.T) {
	level := model.NewLevel(1, []*pb.Testtube{
		model.NewTesttube(4, []pb.Color{pb.Color_RED, pb.Color_GREEN, pb.Color_RED, pb.Color_GREEN}),
		model.NewTesttube(4, []pb.Color{pb.Color_RED, pb.Color_GREEN, pb.Color_RED, pb.Color_GREEN}),
	})
	levelPlay := model.NewLevelPlay(level)

	Solve(levelPlay)
	if !levelPlay.GetCurrentState().GetWon() {
		t.Fatalf(`Game not won. Level: %v`, levelPlay.GetCurrentState().String())
	}
}
