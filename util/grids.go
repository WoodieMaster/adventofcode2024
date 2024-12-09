package util

func DoubleSliceInbound[T any](grid [][]T, col int, row int) bool {
	return row >= 0 && row < len(grid) && col >= 0 && col < len(grid[row])
}

func WithinGrid(gridMin Vec2, gridMax Vec2, pos Vec2) bool {
	return pos.Col >= gridMin.Col && pos.Col < gridMax.Col && pos.Row >= gridMin.Row && pos.Row < gridMax.Row
}

type Vec2 struct {
	Row int
	Col int
}

func AddVec2(a Vec2, b Vec2) Vec2 {
	return Vec2{Col: a.Col + b.Col, Row: a.Row + b.Row}
}

func SubVec2(a Vec2, b Vec2) Vec2 {
	return Vec2{Col: a.Col - b.Col, Row: a.Row - b.Row}
}

func DstVec2(a Vec2, b Vec2) Vec2 {
	return Vec2{Col: Abs(a.Col - b.Col), Row: Abs(a.Row - b.Row)}
}
