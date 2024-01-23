package level

import (
	"log"
	"testing"
)

func TestReadFile(t *testing.T) {
	filename := "textpb/sample.textproto"
	level, err := GetLevelFromFile(filename)
	if err != nil {
		log.Printf("Err: %v", err)
	}
	log.Printf("Level :\n%v", level)
}
