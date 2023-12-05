package main

import (
	"fmt"
	"log"
)

func main() {
	tt := Testtube{4, []string{"red", "green"}}
	fmt.Println(tt)
	err := tt.AddColor("green")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tt)
}
