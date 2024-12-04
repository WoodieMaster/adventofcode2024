package day_4_2

import (
	"adventOfCode/util"
	"fmt"
	"strings"
)

const targetStr = "MAS"

func validXMAS(grid []string, row int, col int) bool {
	util.Assert(grid[row][col] == targetStr[1], "should be "+string(targetStr[1]))

	if row > 0 && row < len(grid)-1 && col > 0 && col < len(grid[row])-1 &&
		(grid[row-1][col-1] == targetStr[0] && grid[row+1][col+1] == targetStr[2] ||
			grid[row-1][col-1] == targetStr[2] && grid[row+1][col+1] == targetStr[0]) &&
		(grid[row-1][col+1] == targetStr[0] && grid[row+1][col-1] == targetStr[2] ||
			grid[row-1][col+1] == targetStr[2] && grid[row+1][col-1] == targetStr[0]) {

		return true
	}
	return false
}

func Main() {
	data := util.LoadFile(4)

	grid := strings.Split(data, "\n")

	result := 0

	for row, line := range grid {
		for col, char := range line {
			if char == int32(targetStr[1]) && validXMAS(grid, row, col) {
				result++
			}
		}
	}

	fmt.Println(result)
}
