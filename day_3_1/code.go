package day_3_1

import (
	"adventOfCode/util"
	"fmt"
	"regexp"
	"strconv"
)

func Main() {
	data := util.LoadFile(3)
	regex := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	matches := regex.FindAllStringSubmatch(data, -1)

	result := 0
	for _, match := range matches {
		result += util.Must(strconv.Atoi(match[1])) * util.Must(strconv.Atoi(match[2]))
	}

	fmt.Println(result)
}
