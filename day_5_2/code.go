package day_5_2

import (
	"adventOfCode/util"
	"fmt"
	"strconv"
	"strings"
)

func parseRule(line string) (int, int) {
	parts := strings.Split(line, "|")

	return util.Must(strconv.Atoi(parts[0])), util.Must(strconv.Atoi(parts[1]))
}

func parseChanges(line string) []int {
	nums := strings.Split(line, ",")

	var result []int

	for _, num := range nums {
		result = append(result, util.Must(strconv.Atoi(num)))
	}

	return result
}

func parseLines(text string) (rules map[int][]int, changes [][]int) {
	lines := strings.Split(text, "\n")
	i := 0

	rules = map[int][]int{}
	for ; i < len(lines); i++ {
		line := lines[i]
		line = strings.TrimSpace(line)

		// double empty line separates rules from the pages
		if len(line) == 0 {
			break
		}

		key, val := parseRule(line)

		rule, ok := rules[key]

		if !ok {
			rules[key] = []int{val}
		} else {
			rules[key] = append(rule, val)
		}
	}

	changes = [][]int{}

	for i++; i < len(lines); i++ {
		line := lines[i]
		line = strings.TrimSpace(line)

		changes = append(changes, parseChanges(line))
	}

	return
}

func validateChanges(rules map[int][]int, changes []int) int {
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

func Main() {
	text := util.LoadFile(5)

	rules, changes := parseLines(text)

	result := 0
	parseLines(text)
	for _, change := range changes {
		result += validateChanges(rules, change)
	}

	fmt.Println(result)
}
