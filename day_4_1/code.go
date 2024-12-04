package day_4_1

import (
	"adventOfCode/util"
	"fmt"
	"strings"
)

const targetStr = "XMAS"

func countFromX(grid []string, row int, col int) int {
	count := 0

	util.Assert(grid[row][col] == targetStr[0], "should be "+string(targetStr[0]))

	for rowMul := -1; rowMul <= 1; rowMul++ {
		for colMul := -1; colMul <= 1; colMul++ {
			valid := true
			for idx, char := range targetStr {
				if idx == 0 {
					continue
				}

				rowPos := row + idx*rowMul
				colPos := col + idx*colMul

				if rowPos < 0 || rowPos >= len(grid) || colPos < 0 || colPos >= len(grid[rowPos]) || int32(grid[rowPos][colPos]) != char {
					valid = false
					break
				}
			}
			if valid {
				count++
			}
		}
	}

	return count
}

func Main() {
	data := util.LoadFile(4)

	lines := strings.Split(data, "\n")

	result := 0

	for row, line := range lines {
		for col, char := range line {
			if char == int32(targetStr[0]) {
				result += countFromX(lines, row, col)
			}
		}
	}

	fmt.Println(result)
}
