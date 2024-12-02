package day_2_1

import (
	"adventOfCode/util"
	"fmt"
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
	diff := 0
	prev := line[0]
	for _, number := range line[1:] {
		newDiff := number - prev

		absDiff := util.Abs(int64(newDiff))

		if absDiff < 1 || absDiff > 3 || (newDiff < 0 && diff > 0) || (newDiff > 0 && diff < 0) {
			return false
		}

		diff = newDiff
		prev = number
	}
	return true
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
