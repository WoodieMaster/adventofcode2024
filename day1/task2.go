package day1

import (
	"fmt"
)

func Task2() {
	lines := loadData()

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
