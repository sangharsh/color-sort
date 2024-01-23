package level

import (
	"log"
	"os"

	"google.golang.org/protobuf/encoding/prototext"

	pb "github.com/sangharsh/color-sort/gen/modelpb"
)

func GetLevelFromFile(filename string) (*pb.LevelState, error) {
	in, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalln("Error reading file:", err)
		return nil, err
	}
	level := &pb.LevelState{}
	if err := prototext.Unmarshal(in, level); err != nil {
		log.Fatalln("Failed to parse level:", err)
	}
	return level, nil
}
