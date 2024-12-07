package day5

import (
	"adventOfCode/util"
	"fmt"
)

func validateChanges1(rules map[int][]int, changes []int) int {
	seenNums := map[int]bool{}

	for _, num := range changes {
		rule, ok := rules[num]
		if ok {
			for _, ruleNum := range rule {
				if seenNums[ruleNum] {
					return 0
				}
			}
		}
		seenNums[num] = true
	}

	return changes[len(changes)/2]
}

func Task1() {
	text := util.LoadFile(5)

	rules, changes := parseLines(text)

	result := 0
	parseLines(text)
	for _, change := range changes {
		result += validateChanges1(rules, change)
	}

	fmt.Println(result)
}
