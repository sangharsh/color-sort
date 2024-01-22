package solver

import (
	"log"
	"math/rand"
	"time"

	pb "github.com/sangharsh/color-sort/gen/modelpb"
	"github.com/sangharsh/color-sort/model"
)

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
				log.Printf("Move: src:%v, dst:%v\n", src, dst)
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

func RandomSolve(levelPlay *pb.LevelPlay) {

	// !levelPlay.GetCurrentState().GetWon()
	for i := 0; i < 50; i++ {
		oneMove(levelPlay)
	}
}
