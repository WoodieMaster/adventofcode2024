package day6

import (
	"adventOfCode/util"
	"fmt"
	"strconv"
	"strings"
)

const (
	Empty    uint8 = '.'
	Obstacle uint8 = '#'
	Start    uint8 = '^'
	Visited  uint8 = 'X'
)

const (
	Up    Direction = 0
	Right Direction = 1
	Down  Direction = 2
	Left  Direction = 3
)

func turnRight(dir Direction) Direction {
	return (dir + 1) % 4
}

func printGrid(grid Grid) {
	for _, line := range grid {
		for _, ch := range line {
			fmt.Print(string(ch))
		}
		fmt.Println()
	}
}

func NextMoveDir(grid Grid, position Point, direction Direction) (Direction, bool) {
	nextPos := addPoint(position, dirToPosChange(direction))

	if !util.SliceGridInbound(grid, nextPos.Col, nextPos.Row) {
		return 0, false
	}

	if grid[nextPos.Row][nextPos.Col] != Obstacle {
		return direction, true
	}
	return NextMoveDir(grid, position, turnRight(direction))
}

func dirToPosChange(dir Direction) Point {
	switch dir {
	case Up:
		return Point{Row: -1, Col: 0}
	case Right:
		return Point{Row: 0, Col: 1}
	case Down:
		return Point{Row: 1, Col: 0}
	case Left:
		return Point{Row: 0, Col: -1}
	}
	panic("Invalid dir " + strconv.Itoa(int(dir)))
}

type Point struct {
	Row int
	Col int
}

func addPoint(a Point, b Point) Point {
	return Point{Col: a.Col + b.Col, Row: a.Row + b.Row}
}

type Direction uint8

type Grid [][]uint8

func loadAndParseData() (grid Grid, start Point) {
	data := util.LoadFile(6)

	lines := strings.Split(data, "\n")

	for rowIdx, line := range lines {
		var gridLine []uint8
		for colIdx, ch := range line {
			cellType := Empty
			switch ch {
			case '#':
				cellType = Obstacle
			case '^':
				cellType = Start
				start = Point{Row: rowIdx, Col: colIdx}
			}

			gridLine = append(gridLine, cellType)
		}
		grid = append(grid, gridLine)
	}

	return
}
