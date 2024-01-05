package level

import (
	"math/rand"

	pb "github.com/sangharsh/color-sort/gen/modelpb"
	"github.com/sangharsh/color-sort/model"
)

var (
	minTubes   = 2
	colorArray = [...]pb.Color{
		pb.Color_RED,
		pb.Color_GREEN,
		pb.Color_BLUE,
		pb.Color_YELLOW,
		pb.Color_GRAY,
		pb.Color_LIME_GREEN,
		pb.Color_VIOLET,
		pb.Color_PINK,
		pb.Color_ORANGE,
		pb.Color_SKY_BLUE,
		pb.Color_LIGHT_GREEN,
		pb.Color_BROWN,
	}
)

func Generate(level int32) *pb.LevelState {
	r := rand.New(rand.NewSource(int64(level)))
	numTubes := minTubes + r.Intn(len(colorArray)-minTubes)
	tubes := []*pb.Testtube{}
	colorDist := [][]int{r.Perm(numTubes), r.Perm(numTubes), r.Perm(numTubes), r.Perm(numTubes)}

	for i := 0; i < numTubes; i++ {
		tubes = append(tubes, model.NewTesttube(4,
			[]pb.Color{
				colorArray[colorDist[0][i]],
				colorArray[colorDist[1][i]],
				colorArray[colorDist[2][i]],
				colorArray[colorDist[3][i]],
			}))
	}
	gl := model.NewLevel(level, tubes)
	return gl
}
