package day2

import (
	"adventOfCode/util"
	"fmt"
)

func validateLine1(line []int) bool {
	diff := 0
	prev := line[0]
	for _, number := range line[1:] {
		newDiff := number - prev

		absDiff := util.Abs(newDiff)

		if absDiff < 1 || absDiff > 3 || (newDiff < 0 && diff > 0) || (newDiff > 0 && diff < 0) {
			return false
		}

		diff = newDiff
		prev = number
	}
	return true
}

func Task1() {
	lines := loadData()

	count := 0
	for _, line := range lines {
		if validateLine1(readLine(line)) {
			count++
		}
	}

	fmt.Println(count)
}
