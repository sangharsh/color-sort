package solver

import (
	pb "github.com/sangharsh/color-sort/gen/modelpb"
	"github.com/sangharsh/color-sort/model"
	"google.golang.org/protobuf/proto"
)

type BFSSolver struct {
	initialState *pb.LevelState
	visited      map[string]bool
	queue        []*pb.LevelPlay
}

func NewBFSSolver(levelState *pb.LevelState) *BFSSolver {
	s := new(BFSSolver)
	s.initialState = levelState
	s.visited = make(map[string]bool)
	s.queue = make([]*pb.LevelPlay, 0)
	return s
}

func (s BFSSolver) Solve() *pb.LevelPlay {
	levelState := s.initialState
	levelPlay := model.NewLevelPlay(levelState)
	clone := proto.Clone(levelPlay).(*pb.LevelPlay)
	s.queue = append(s.queue, clone)
	return s.solve2()
}

func (s BFSSolver) solve2() *pb.LevelPlay {
	movesTried := 0
	statesTried := 0
	for len(s.queue) > 0 {
		levelPlay := s.queue[0]
		s.queue = s.queue[1:]
		AddToVisited(s.visited, levelPlay)

		if levelPlay.GetCurrentState().GetWon() {
			return levelPlay
		}

		tubes := levelPlay.GetCurrentState().GetTubes()
		numTubes := len(tubes)
		var pourRes *pb.PourResponse
		for src := 0; src < numTubes; src++ {
			for dst := 0; dst < numTubes; dst++ {
				movesTried++
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
				} else if IsVisited(s.visited, levelPlay) {
					model.Undo(&pb.UndoRequest{}, levelPlay)
					continue
				} else {
					statesTried++
					clone := proto.Clone(levelPlay).(*pb.LevelPlay)
					s.queue = append(s.queue, clone)
					model.Undo(&pb.UndoRequest{}, levelPlay)
				}
			}
		}
	}
	return nil
}
