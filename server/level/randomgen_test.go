package level

import (
	"fmt"
	"math/rand"
	"testing"

	pb "github.com/sangharsh/color-sort/gen/modelpb"
)

func TestLevelNumber(t *testing.T) {
	level := 5
	gl := Generate(int32(level))
	fmt.Printf("gl: %v\n", gl.Glpb.String())
	if gl.Glpb.GetLevel() != int32(level) {
		t.Fatalf("Level not matching")
	}
}

func TestTubesAndColors(t *testing.T) {
	gl := Generate(rand.Int31())
	tubes := gl.Glpb.GetTubes()
	colorCountMap := make(map[pb.Color]int)

	for i := 0; i < len(tubes); i++ {
		colors := tubes[i].GetColors()
		for j := 0; j < len(colors); j++ {
			colorCountMap[colors[j]] = colorCountMap[colors[j]] + 1
		}
	}

	if len(tubes) != len(colorCountMap)+2 {
		t.Fatalf("Number of tubes and colors not matching. \nLevel: %v", gl)
	}

	for color, count := range colorCountMap {
		if count != 4 {
			t.Fatalf("Count is not 4 for color %v, \nLevel: %v", color, gl)
		}
	}

}
