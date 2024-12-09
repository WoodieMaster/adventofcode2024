package day1

import (
	"adventOfCode/util"
	"fmt"
	"sort"
)

func Task1() {
	lines := loadData()
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
		diff += int64(util.Abs(list1[i] - list2[i]))
	}

	fmt.Println(diff)
}
