package solver

import (
	pb "github.com/sangharsh/color-sort/gen/modelpb"
	"github.com/sangharsh/color-sort/model"
)

type DFSSolver struct {
	initialState *pb.LevelState
	visited      map[string]bool
}

func NewDFSSolver(levelState *pb.LevelState) *DFSSolver {
	s := new(DFSSolver)
	s.initialState = levelState
	s.visited = make(map[string]bool)
	return s
}

func (s DFSSolver) Solve() *pb.LevelPlay {
	levelState := s.initialState
	levelPlay := model.NewLevelPlay(levelState)
	s.solve2(levelPlay)
	return levelPlay
}

func (s DFSSolver) solve2(levelPlay *pb.LevelPlay) {
	if levelPlay.GetCurrentState().GetWon() {
		return
	}
	if IsVisited(s.visited, levelPlay) {
		return
	} else {
		AddToVisited(s.visited, levelPlay)
	}

	tubes := levelPlay.GetCurrentState().GetTubes()
	numTubes := len(tubes)
	var pourRes *pb.PourResponse
	for src := 0; src < numTubes; src++ {
		for dst := 0; dst < numTubes; dst++ {
			if src == dst {
				continue
			}
			pourReq := pb.PourRequest{
				Src: int32(src),
				Dst: int32(dst),
			}
			pourRes = model.Pour(&pourReq, levelPlay)
			if pourRes.GetErr() != "" {
				continue
			} else {
				s.solve2(levelPlay)
				if levelPlay.GetCurrentState().GetWon() {
					return
				}
				model.Undo(&pb.UndoRequest{}, levelPlay)
			}
		}
	}
}
