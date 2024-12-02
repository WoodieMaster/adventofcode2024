package day_2_2

import (
	"adventOfCode/util"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func readLine(line string) []int {
	var numbers []int

	for _, numberStr := range strings.Split(line, " ") {
		readyStr := strings.Trim(numberStr, " \r\n\t")
		number := util.Must(strconv.ParseInt(readyStr, 10, 32))

		numbers = append(numbers, int(number))
	}

	return numbers
}

func validateLine(line []int) bool {
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

func Main() {
	data := util.LoadFile(2)

	lines := strings.Split(data, "\n")

	count := 0
	for _, line := range lines {
		if validateLine(readLine(line)) {
			count++
		}
	}

	fmt.Println(count)
}
