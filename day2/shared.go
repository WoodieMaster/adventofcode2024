package day2

import (
	"adventOfCode/util"
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

func loadData() []string {
	data := util.LoadFile(2)

	return strings.Split(data, "\n")
}
