package level

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/sangharsh/color-sort/model"
)

func TestLevelNumber(t *testing.T) {
	levelId := 5
	level := Generate(int32(levelId))
	fmt.Printf("gl: %v\n", level.String())
	if level.GetId() != int32(levelId) {
		t.Fatalf("Level not matching")
	}
}

func TestTubesAndColors(t *testing.T) {
	level := Generate(rand.Int31())
	ok, err := model.HasValidColorAndTubes(level)
	if !ok || err != nil {
		t.Fatalf("Number of tubes or colors not matching. \nLevel: %v", level)
	}
}
