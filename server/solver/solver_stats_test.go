package solver

import (
	"testing"

	"github.com/sangharsh/color-sort/level"
)

func TestGreedySolverStats(t *testing.T) {

	t.Logf("Level\tColors\tMoves\tPours\tUndos\tExMoves\tEffy")
	for i := 0; i < 1000; i++ {
		level := level.Generate(int32(i))
		solver := NewGreedySolver(level)
		levelPlay := solver.Solve()

		if !levelPlay.GetCurrentState().GetWon() {
			t.Fatalf(`Game not won. Level: %v`, levelPlay.GetCurrentState().String())
		} else {
			levelId := levelPlay.CurrentState.Id
			numColors := len(level.GetTubes()) - 2
			numMoves := len(levelPlay.Moves)
			t.Logf("%v\t%v\t%v\t%v\t%v\t%v\t%v",
				levelId, numColors, numMoves,
				solver.Stats().pourCount, solver.Stats().undoCount,
				numMoves-3*numColors, float32(numMoves)/float32(numColors))
		}
	}
}
