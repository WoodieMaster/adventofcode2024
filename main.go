package main

import (
	"adventOfCode/day_1_1"
	"adventOfCode/day_1_2"
	"adventOfCode/day_2_1"
	"adventOfCode/day_2_2"
	"adventOfCode/day_3_1"
	"adventOfCode/day_3_2"
	"adventOfCode/day_4_1"
	"adventOfCode/day_4_2"
	"adventOfCode/day_5_1"
	"adventOfCode/day_5_2"
	"log"
	"os"
)

var puzzles = map[string]func(){
	"1_1": day_1_1.Main,
	"1_2": day_1_2.Main,
	"2_1": day_2_1.Main,
	"2_2": day_2_2.Main,
	"3_1": day_3_1.Main,
	"3_2": day_3_2.Main,
	"4_1": day_4_1.Main,
	"4_2": day_4_2.Main,
	"5_1": day_5_1.Main,
	"5_2": day_5_2.Main,
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
