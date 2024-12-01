package main

import (
	"adventOfCode/day_1_1"
	"adventOfCode/day_1_2"
	"log"
	"os"
)

var puzzles = map[string]func(){
	"1_1": day_1_1.Main,
	"1_2": day_1_2.Main,
}

func main() {

	if len(os.Args) < 2 {
		log.Fatal("usage: code <day> [-s]")
	}

	fn, ok := puzzles[os.Args[1]]

	if !ok {
		log.Fatal("Invalid day")
	}

	fn()
}
