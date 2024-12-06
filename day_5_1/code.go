package day_5_1

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
