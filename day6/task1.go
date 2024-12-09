package day6

import (
	"adventOfCode/util"
	"fmt"
)

func traceGuardSteps(grid Grid, start util.Vec2) int {
	visitedPoints := map[util.Vec2]bool{}
	currentPos := start
	currentDir := Up

	for {
		dir, inside := NextMoveDir(grid, currentPos, currentDir)
		if !inside {
			break
		}
		currentDir = dir
		currentPos = util.AddVec2(currentPos, dirToPosChange(currentDir))
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
