package db

import (
	"log"

	pb "github.com/sangharsh/color-sort/gen/modelpb"
)

var (
	userIdLevePlayMap = make(map[string]*pb.LevelPlay)
)

func Get(userId string) *pb.LevelPlay {
	return userIdLevePlayMap[userId]
}

func Set(userId string, levelPlay *pb.LevelPlay) (bool, error) {
	userIdLevePlayMap[userId] = levelPlay
	log.Printf("Map: %v", userIdLevePlayMap)
	return true, nil
}
