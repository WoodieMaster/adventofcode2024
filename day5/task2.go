package day5

import (
	"adventOfCode/util"
	"fmt"
)

func validateChanges2(rules map[int][]int, changes []int) int {
	seenNums := map[int]int{}

	outOfOrder := false

	for idx := 0; idx < len(changes); idx++ {
		num := changes[idx]
		seenNums[num] = idx

		rule, ok := rules[num]
		if !ok {
			continue
		}
		moveIdx := idx
		for _, ruleNum := range rule {
			ruleIdx, ok := seenNums[ruleNum]
			if !ok || // rule number not found yet -> cannot be in front
				ruleIdx > idx || // rule number already after current number
				ruleIdx >= moveIdx { // moving the number to current moveIdx already satisfies this rule
				continue
			}
			moveIdx = ruleIdx
		}

		if moveIdx < idx {
			changes[moveIdx], changes[idx] = changes[idx], changes[moveIdx]

			seenNums[changes[moveIdx]] = moveIdx
			seenNums[changes[idx]] = idx

			idx = moveIdx + 1
			outOfOrder = true
		}
	}

	if !outOfOrder {
		return 0
	}

	return changes[len(changes)/2]
}

func Task2() {
	text := util.LoadFile(5)

	rules, changes := parseLines(text)

	result := 0
	parseLines(text)
	for _, change := range changes {
		result += validateChanges2(rules, change)
	}

	fmt.Println(result)
}
