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

func NextMoveDir(grid Grid, position util.Vec2, direction Direction) (Direction, bool) {
	nextPos := util.AddVec2(position, dirToPosChange(direction))

	if !util.DoubleSliceInbound(grid, nextPos.Col, nextPos.Row) {
		return 0, false
	}

	if grid[nextPos.Row][nextPos.Col] != Obstacle {
		return direction, true
	}
	return NextMoveDir(grid, position, turnRight(direction))
}

func dirToPosChange(dir Direction) util.Vec2 {
	switch dir {
	case Up:
		return util.Vec2{Row: -1, Col: 0}
	case Right:
		return util.Vec2{Row: 0, Col: 1}
	case Down:
		return util.Vec2{Row: 1, Col: 0}
	case Left:
		return util.Vec2{Row: 0, Col: -1}
	}
	panic("Invalid dir " + strconv.Itoa(int(dir)))
}

type Direction uint8

type Grid [][]uint8

func loadAndParseData() (grid Grid, start util.Vec2) {
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
				start = util.Vec2{Row: rowIdx, Col: colIdx}
			}

			gridLine = append(gridLine, cellType)
		}
		grid = append(grid, gridLine)
	}

	return
}
