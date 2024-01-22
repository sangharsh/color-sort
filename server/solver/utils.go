package solver

import (
	"slices"
	"strconv"
	"strings"

	pb "github.com/sangharsh/color-sort/gen/modelpb"
	"google.golang.org/protobuf/proto"
)

func numColors(tt *pb.Testtube) int {
	count := 0
	colorMap := make(map[pb.Color]bool)
	for _, color := range tt.GetColors() {
		if _, ok := colorMap[color]; !ok {
			colorMap[color] = true
			count++
		}
	}
	return count
}
func Entropy(level *pb.LevelState) int {
	// Define as number of colors in each tube
	// Pure tubes?
	score := 0
	for _, tt := range level.GetTubes() {
		score += numColors(tt)
	}
	return score
}

func AddToVisited(visited map[string]bool, levelPlay *pb.LevelPlay) {
	clone := proto.Clone(levelPlay).(*pb.LevelPlay)
	tubes := clone.GetCurrentState().GetTubes()
	slices.SortFunc(tubes, TesttubeSortFn)
	visited[clone.GetCurrentState().String()] = true
}

func IsVisited(visited map[string]bool, levelPlay *pb.LevelPlay) bool {
	clone := proto.Clone(levelPlay).(*pb.LevelPlay)
	tubes := clone.GetCurrentState().GetTubes()
	slices.SortFunc(tubes, TesttubeSortFn)
	return visited[clone.GetCurrentState().String()]
}

func MovesString(levelPlay *pb.LevelPlay) string {
	var sb strings.Builder

	for _, move := range levelPlay.GetMoves() {
		sb.WriteString(strconv.Itoa(int(move.GetSrc())) + " -> " + strconv.Itoa(int(move.GetDst())) + ", ")
	}
	return sb.String()
}

func TesttubeSortFn(a, b *pb.Testtube) int {
	aCols := a.GetColors()
	bCols := b.GetColors()
	if len(aCols) != len(bCols) {
		return len(aCols) * (len(aCols) - len(bCols))
	} else {
		for i := 0; i < len(aCols); i++ {
			if aCols[i] != bCols[i] {
				return (len(aCols) - i) * int(aCols[i].Number()-bCols[i].Number())
			}
		}
	}
	return 0
}
