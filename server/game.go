package main

import (
	"fmt"
	"log"
	"github.com/sangharsh/color-sort/model"
)

func main() {
	tt := model.NewTesttube(4, []string{"red", "green"})
	tt2 := model.NewTesttube(4, []string{"red", "green"})
	gameLevel := model.NewGameLevel(1, []model.Testtube{*tt, *tt2})
	fmt.Println(gameLevel)
	_, err := gameLevel.Pour(0, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(gameLevel)
	fmt.Println("Won: ", gameLevel.Won())
}
