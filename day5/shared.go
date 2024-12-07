package day5

import (
	"adventOfCode/util"
	"strconv"
	"strings"
)

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

func parseChanges(line string) []int {
	nums := strings.Split(line, ",")

	var result []int

	for _, num := range nums {
		result = append(result, util.Must(strconv.Atoi(num)))
	}

	return result
}

func parseRule(line string) (int, int) {
	parts := strings.Split(line, "|")

	return util.Must(strconv.Atoi(parts[0])), util.Must(strconv.Atoi(parts[1]))
}
