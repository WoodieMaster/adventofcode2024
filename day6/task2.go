package day6

import (
	"fmt"
	"maps"
	"slices"
)

func addStateToVisited(visitedPoints VisitedRecord, currentPos Point, currentDir Direction) bool {
	dirs, ok := visitedPoints[currentPos]

	if !ok {
		visitedPoints[currentPos] = []Direction{currentDir}
		return true
	}

	if !slices.Contains(dirs, currentDir) {
		visitedPoints[currentPos] = append(dirs, currentDir)
		return true
	}

	return false // Point already visited with the same direction
}

type VisitedRecord map[Point][]Direction

func cloneVisited(src VisitedRecord) VisitedRecord {
	dst := VisitedRecord{}
	for k, v := range src {
		dst[k] = []Direction{}
		copy(dst[k], v)
	}
	maps.Copy(dst, src)
	return dst
}

func checkForLoop(grid Grid, alreadyVisited VisitedRecord, startPos Point, startDir Direction) bool {
	visitedPoints := cloneVisited(alreadyVisited)
	currentPos := startPos
	currentDir := startDir

	for {
		dir, inside := NextMoveDir(grid, currentPos, currentDir)
		if !inside {
			break
		}
		currentDir = dir
		currentPos = addPoint(currentPos, dirToPosChange(currentDir))

		if !addStateToVisited(visitedPoints, currentPos, currentDir) {
			return true
		}
	}

	return false
}

func FindPossibleLoops(grid Grid, startPos Point) (result []Point) {
	visitedPoints := VisitedRecord{}

	currentPos := startPos
	currentDir := Up

	for {
		addStateToVisited(visitedPoints, currentPos, currentDir)

		// normal move code
		nextDir, inside := NextMoveDir(grid, currentPos, currentDir)
		if !inside {
			break
		}
		nextPos := addPoint(currentPos, dirToPosChange(nextDir))

		if grid[nextPos.Row][nextPos.Col] == Empty {
			grid[nextPos.Row][nextPos.Col] = Visited
		}

		if _, exists := visitedPoints[nextPos]; !exists {
			// add obstacle
			grid[nextPos.Row][nextPos.Col] = Obstacle

			// add obstaclePosition to result if it created a loop
			if checkForLoop(grid, visitedPoints, currentPos, currentDir) {
				result = append(result, nextPos)
			}

			// reset grid
			grid[nextPos.Row][nextPos.Col] = Empty
		}

		// move
		currentDir = nextDir
		currentPos = nextPos
	}

	return
}

func Task2() {
	result := FindPossibleLoops(loadAndParseData())

	fmt.Println(result)

	fmt.Println(len(result))
}
