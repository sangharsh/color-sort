package solver

import (
	"math/rand"
	"time"

	pb "github.com/sangharsh/color-sort/gen/modelpb"
	"github.com/sangharsh/color-sort/model"
)

type RandomSolver struct {
	initialState *pb.LevelState
}

func NewRandomSolver(levelState *pb.LevelState) *RandomSolver {
	s := new(RandomSolver)
	s.initialState = levelState
	return s
}

func oneMove(levelPlay *pb.LevelPlay) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	tubes := levelPlay.GetCurrentState().GetTubes()
	numTubes := len(tubes)
	var pourRes *pb.PourResponse
	for src := 0; src < numTubes; src++ {
		dstOptions := r.Perm(numTubes)
		for _, dst := range dstOptions {
			if src == dst {
				continue
			}
			pourReq := pb.PourRequest{
				Src: int32(src),
				Dst: int32(dst),
			}
			pourRes = model.Pour(&pourReq, levelPlay)
			if pourRes.GetErr() == "" {
				break
			}
		}
		if pourRes.GetErr() == "" {
			break
		}
	}
	if pourRes.GetErr() == "" {
		return
	} else {
		model.Undo(&pb.UndoRequest{}, levelPlay)
	}
}

func (s RandomSolver) Solve() *pb.LevelPlay {
	levelPlay := model.NewLevelPlay(s.initialState)
	for i := 0; i < 50; i++ {
		oneMove(levelPlay)
	}
	return levelPlay
}
