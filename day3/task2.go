package day3

import (
	"adventOfCode/util"
	"fmt"
	"regexp"
	"strconv"
)

const mulRegex = `(mul)\((\d+),(\d+)\)`
const conditionRegex = `(do(n't)?)\(\)`

func Task2() {
	data := util.LoadFile(3)
	regex := regexp.MustCompile(mulRegex + "|" + conditionRegex)

	matches := regex.FindAllStringSubmatch(data, -1)

	result := 0
	mulEnabled := true
	for _, match := range matches {
		if match[1] == "mul" {
			if mulEnabled {
				result += util.Must(strconv.Atoi(match[2])) * util.Must(strconv.Atoi(match[3]))
			}
		} else {
			mulEnabled = match[4] == "do"
		}
	}

	fmt.Println(result)
}
