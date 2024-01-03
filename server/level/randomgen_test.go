package level

import (
	"fmt"
	"testing"
)

func TestLevelGen(t *testing.T) {
	level := 5
	gl := Generate(int32(level))
	fmt.Printf("gl: %v\n", gl.Glpb.String())
	if gl.Glpb.GetLevel() != int32(level) {
		t.Fatalf("Level not matching")
	}
}
