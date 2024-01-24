package solver

import (
	"slices"

	pb "github.com/sangharsh/color-sort/gen/modelpb"
	"github.com/sangharsh/color-sort/model"
)

type GreedySolver struct {
	initialState *pb.LevelState
	visited      map[string]bool
	stats        *SolverStat
}

func NewGreedySolver(levelState *pb.LevelState) *GreedySolver {
	s := new(GreedySolver)
	s.initialState = levelState
	s.visited = make(map[string]bool)
	s.stats = new(SolverStat)
	return s
}

func (s GreedySolver) Stats() *SolverStat {
	return s.stats
}

func (s GreedySolver) Solve() *pb.LevelPlay {
	levelState := s.initialState
	levelPlay := model.NewLevelPlay(levelState)
	s.solve2(levelPlay)
	return levelPlay
}

func getScore(src, dst *pb.Testtube) int {
	if numColors(src) == 1 && numColors(dst) == 0 {
		// Avoid move from pure to empty
		return -2
	}
	if numColors(dst) == 1 {
		// Prefer move to pure
		return 3
	}
	if numColors(src) == 1 && int(dst.GetSize())-len(dst.GetColors()) >= len(src.GetColors()) {
		// Prefer creating pure
		return 2
	}
	numSrcColor := 1
	srcCols := src.GetColors()
	for i := 1; i < len(srcCols); i++ {
		if srcCols[len(srcCols)-1] == srcCols[len(srcCols)-1-i] {
			numSrcColor++
		}
	}
	if numSrcColor > int(dst.GetSize())-len(dst.GetColors()) {
		return -1
	}
	return 1
}

type MovePotential struct {
	src       int
	dst       int
	potential int
}

func (s GreedySolver) solve2(levelPlay *pb.LevelPlay) {
	if levelPlay.GetCurrentState().GetWon() {
		return
	}
	if IsVisited(s.visited, levelPlay) {
		return
	} else {
		AddToVisited(s.visited, levelPlay)
	}

	tubes := levelPlay.GetCurrentState().GetTubes()

	potentialMoves := getPotentialMoves(tubes)

	for _, move := range potentialMoves {
		pourReq := pb.PourRequest{
			Src: int32(move.src),
			Dst: int32(move.dst),
		}
		pourRes := model.Pour(&pourReq, levelPlay)
		if pourRes.GetErr() != "" {
			continue
		} else {
			s.stats.pourCount += 1
			s.solve2(levelPlay)
			if levelPlay.GetCurrentState().GetWon() {
				return
			}
			model.Undo(&pb.UndoRequest{}, levelPlay)
			s.stats.undoCount += 1
		}
	}
}

func getPotentialMoves(tubes []*pb.Testtube) []MovePotential {
	var potentialMoves []MovePotential
	for srcIdx, srcTube := range tubes {
		for dstIdx, dstTube := range tubes {
			if srcIdx == dstIdx {
				// Same tube
				continue
			}
			srcCols := srcTube.GetColors()
			dstCols := dstTube.GetColors()
			if len(srcCols) == 0 || len(dstCols) == int(dstTube.GetSize()) {
				// Empty src or full dst
				continue
			}
			if len(dstCols) != 0 && srcCols[len(srcCols)-1] != dstCols[len(dstCols)-1] {
				// Top colors non matching
				continue
			}
			potentialMoves = append(potentialMoves, MovePotential{srcIdx, dstIdx, getScore(srcTube, dstTube)})
		}
	}
	slices.SortFunc(potentialMoves, func(a, b MovePotential) int { return b.potential - a.potential })
	return potentialMoves
}
