package main

import (
	"adventOfCode/day1"
	"adventOfCode/day2"
	"adventOfCode/day3"
	"adventOfCode/day4"
	"adventOfCode/day5"
	"adventOfCode/day6"
	"fmt"
	"log"
	"os"
)

var puzzles = map[string]func(){
	"1_1": day1.Task1,
	"1_2": day1.Task2,
	"2_1": day2.Task1,
	"2_2": day2.Task2,
	"3_1": day3.Task1,
	"3_2": day3.Task2,
	"4_1": day4.Task1,
	"4_2": day4.Task2,
	"5_1": day5.Task1,
	"5_2": day5.Task2,
	"6_1": day6.Task1,
	"6_2": day6.Task2,
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: code <task> [-s]")
	}

	fn, ok := puzzles[os.Args[1]]

	if !ok {
		log.Fatal("Invalid task")
	}

	fmt.Println("Starting code for task ", os.Args[1])
	fn()
}
