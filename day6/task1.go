package day6

import (
	"fmt"
)

func traceGuardSteps(grid Grid, start Point) int {
	visitedPoints := map[Point]bool{}
	currentPos := start
	currentDir := Up

	for {
		dir, inside := NextMoveDir(grid, currentPos, currentDir)
		if !inside {
			break
		}
		currentDir = dir
		currentPos = addPoint(currentPos, dirToPosChange(currentDir))
		visitedPoints[currentPos] = true
		if grid[currentPos.Row][currentPos.Col] == Empty {
			grid[currentPos.Row][currentPos.Col] = Visited
		}
	}

	printGrid(grid)
	return len(visitedPoints)
}

func Task1() {
	result := traceGuardSteps(loadAndParseData())

	fmt.Println(result)
}
