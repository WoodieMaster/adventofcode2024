package day_1_2

import (
	"adventOfCode/util"
	"fmt"
	"strconv"
	"strings"
)

func readLine(line string) (a int, b int) {
	completeFirst := false
	numberStr := ""

	for _, code := range line {
		if code >= '0' && code <= '9' {
			numberStr += string(code)
		} else {
			if len(numberStr) == 0 {
				continue
			}
			number := int(util.Must(strconv.ParseInt(numberStr, 10, 32)))
			if completeFirst {
				b = number
				return
			}
			a = number
			completeFirst = true
			numberStr = ""
		}
	}

	if len(numberStr) != 0 {
		panic("should not happen")
	}
	return
}

func Main() {
	data := util.LoadFile(1)

	lines := strings.Split(data, "\n")

	var list1 []int
	map2 := map[int]int{}

	for _, line := range lines {
		n1, n2 := readLine(line)
		list1 = append(list1, n1)

		val, ok := map2[n2]

		if !ok {
			val = 0
		}
		map2[n2] = val + 1
	}

	result := int64(0)

	for i := 0; i < len(list1); i++ {
		num := list1[i]
		result += int64(num * map2[num])
	}

	fmt.Println(result)
}
