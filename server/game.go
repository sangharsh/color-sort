package main

import (
	"fmt"
	"log"

	"github.com/sangharsh/color-sort/model"
)

func main() {
	tt := model.NewTesttube(4, []model.Color{model.Color_RED, model.Color_GREEN, model.Color_RED, model.Color_GREEN})
	tt2 := model.NewTesttube(4, []model.Color{model.Color_RED, model.Color_GREEN, model.Color_RED, model.Color_GREEN})
	gameLevel := model.NewGameLevel(1, []*model.Testtube{tt, tt2})
	fmt.Println(gameLevel)
	fmt.Println("Won: ", gameLevel.Won())
	solve := [][]int{{0, 2}, {0, 3}, {0, 2}, {0, 3}, {1, 2}, {1, 3}, {1, 2}, {1, 3}}

	for _, move := range solve {
		_, err := gameLevel.Pour(move[0], move[1])
		if err != nil {
			fmt.Println(gameLevel)
			log.Fatal(err)
		}
	}

	fmt.Println(gameLevel)
	fmt.Println("Won: ", gameLevel.Won())
}
