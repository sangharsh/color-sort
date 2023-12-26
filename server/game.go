package main

import (
	"fmt"
	"log"
	"github.com/sangharsh/color-sort/model"
)

func main() {
	tt := model.NewTesttube(4, []string{"red", "green", "red", "green"})
	tt2 := model.NewTesttube(4, []string{"red", "green", "red", "green"})
	gameLevel := model.NewGameLevel(1, []model.Testtube{*tt, *tt2})
	fmt.Println(gameLevel)
	solve := [][]int{{0,2}, {0,3}, {0,2}, {0,3}, {1,2}, {1,3}, {1,2}, {1,3}}

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
