package solver

import pb "github.com/sangharsh/color-sort/gen/modelpb"

type Solver interface {
	Solve() *pb.LevelPlay
}
