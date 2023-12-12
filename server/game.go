package main

import (
	"fmt"
	"log"
)

func main() {
	tt := Testtube{4, []string{"red", "green"}}
	tt2 := Testtube{4, []string{"red", "green"}}
	level := GameLevel{1, []Testtube{tt, tt2}}
	fmt.Println(level)
	_, err := level.Pour(0, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(level)
}
