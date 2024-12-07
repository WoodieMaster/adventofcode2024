package day1

import (
	"adventOfCode/util"
	"strconv"
	"strings"
)

func readLine(line string) (a int, b int) {
	completeFirst := false
	numberStr := ""

	for _, code := range line {
		if code >= '0' && code <= '9' {
			numberStr += string(code)
		} else {
			if len(numberStr) == 0 {
				continue
			}
			number := int(util.Must(strconv.ParseInt(numberStr, 10, 32)))
			if completeFirst {
				b = number
				return
			}
			a = number
			completeFirst = true
			numberStr = ""
		}
	}

	if len(numberStr) != 0 {
		panic("should not happen")
	}
	return
}

func loadData() []string {
	data := util.LoadFile(1)

	return strings.Split(data, "\n")
}
