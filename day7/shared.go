package day7

import (
	"adventOfCode/util"
	"regexp"
	"strconv"
	"strings"
)

type Operation struct {
	Result  int
	Numbers []int
}

const (
	Plus   Operator = "+"
	Mult   Operator = "*"
	Concat Operator = "|"
)

type Operator string

func loadAndParseData() (result []Operation) {
	regexNumberList := util.Must(regexp.Compile(` (\d+)`))
	regexOpResult := util.Must(regexp.Compile(`^\d+`))

	text := util.LoadFile(7)
	lines := strings.Split(text, "\n")

	for _, line := range lines {
		regexResult := regexNumberList.FindAllStringSubmatch(line, -1)

		var numbers []int

		for _, match := range regexResult {
			num := util.Must(strconv.Atoi(match[1]))
			numbers = append(numbers, num)
		}

		opResultStr := regexOpResult.FindString(line)
		opResult := util.Must(strconv.Atoi(opResultStr))

		result = append(result, Operation{
			Result:  opResult,
			Numbers: numbers,
		})
	}

	return
}

func concatNumbers(first int, second int) int {
	result := first
	reducer := second

	for reducer > 0 {
		result *= 10
		reducer /= 10
	}

	return result + second
}

func calculate(a int, b int, operator Operator) int {
	switch operator {
	case Plus:
		return a + b
	case Mult:
		return a * b
	case Concat:
		return concatNumbers(a, b)
	}
	panic("Invalid operator")
}

func hasValidCalculation(operation Operation, validOperators ...Operator) bool {
	for permutation := range util.Permutations(len(operation.Numbers)-1, validOperators...) {
		total := operation.Numbers[0]
		for idx, operator := range permutation {
			total = calculate(total, operation.Numbers[idx+1], operator)
		}
		if total == operation.Result {
			return true
		}
	}
	return false
}
