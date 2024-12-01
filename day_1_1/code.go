package day_1_1

import (
	"adventOfCode/util"
	"fmt"
	"sort"
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
	var list2 []int

	for _, line := range lines {
		n1, n2 := readLine(line)
		list1 = append(list1, n1)
		list2 = append(list2, n2)
	}

	sort.Ints(list1)
	sort.Ints(list2)

	diff := int64(0)

	for i := 0; i < len(list1); i++ {
		//fmt.Println(list1[i], list2[i], util.Abs(int64(list1[i]-list2[i])))
		diff += util.Abs(int64(list1[i] - list2[i]))
	}

	fmt.Println(diff)
}
