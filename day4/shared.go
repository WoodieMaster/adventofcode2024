package day4

import (
	"adventOfCode/util"
	"strings"
)

func loadData() []string {
	data := util.LoadFile(4)

	return strings.Split(data, "\n")
}
