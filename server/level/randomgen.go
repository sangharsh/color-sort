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
	colors := []pb.Color{}
	for i := 0; i < 4; i++ {
		colors = append(colors, colorArray[:numTubes]...)
	}
	r.Shuffle(4*numTubes, func(i, j int) {
		colors[i], colors[j] = colors[j], colors[i]
	})

	for i := 0; i < numTubes; i++ {
		tubes = append(tubes, model.NewTesttube(4,
			[]pb.Color{
				colors[4*i+0],
				colors[4*i+1],
				colors[4*i+2],
				colors[4*i+3],
			}))
	}
	gl := model.NewLevel(level, tubes)
	return gl
}
