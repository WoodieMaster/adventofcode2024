package day2

import (
	"adventOfCode/util"
	"fmt"
	"slices"
)

func validateLine2(line []int) bool {
Outer:
	for skipIdx := range line {
		usedLine := slices.Concat(line[0:skipIdx], line[skipIdx+1:])

		diff := 0
		prev := usedLine[0]

		for _, number := range usedLine[1:] {
			newDiff := number - prev

			absDiff := util.Abs(int64(newDiff))

			if absDiff < 1 || absDiff > 3 || (newDiff < 0 && diff > 0) || (newDiff > 0 && diff < 0) {
				continue Outer
			}

			diff = newDiff
			prev = number
		}
		return true
	}
	return false
}

func Task2() {
	lines := loadData()

	count := 0
	for _, line := range lines {
		if validateLine2(readLine(line)) {
			count++
		}
	}

	fmt.Println(count)
}
