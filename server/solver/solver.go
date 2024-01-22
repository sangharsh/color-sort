package solver

import pb "github.com/sangharsh/color-sort/gen/modelpb"

type Solver interface {
	Initial() *pb.LevelState
	Solve() *pb.LevelPlay
}
