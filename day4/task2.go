package day4

import (
	"adventOfCode/util"
	"fmt"
)

const targetStr2 = "MAS"

func validXMAS(grid []string, row int, col int) bool {
	util.Assert(grid[row][col] == targetStr2[1], "should be "+string(targetStr2[1]))

	if row > 0 && row < len(grid)-1 && col > 0 && col < len(grid[row])-1 &&
		(grid[row-1][col-1] == targetStr2[0] && grid[row+1][col+1] == targetStr2[2] ||
			grid[row-1][col-1] == targetStr2[2] && grid[row+1][col+1] == targetStr2[0]) &&
		(grid[row-1][col+1] == targetStr2[0] && grid[row+1][col-1] == targetStr2[2] ||
			grid[row-1][col+1] == targetStr2[2] && grid[row+1][col-1] == targetStr2[0]) {

		return true
	}
	return false
}

func Task2() {
	grid := loadData()

	result := 0

	for row, line := range grid {
		for col, char := range line {
			if char == int32(targetStr2[1]) && validXMAS(grid, row, col) {
				result++
			}
		}
	}

	fmt.Println(result)
}
