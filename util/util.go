package util

import (
	"os"
	"slices"
	"strconv"
)

func LoadFile(day int) string {
	path := "puzzle_input" + string(os.PathSeparator) + strconv.FormatInt(int64(day), 10)
	args := os.Args

	if slices.Contains(args, "-s") {
		path += "_sample"
	}

	bytes, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	return string(bytes)
}

func Must[T any](data T, err error) T {
	if err != nil {
		panic(err)
	}

	return data
}

func SliceGridInbound[T any](grid [][]T, col int, row int) bool {
	return row >= 0 && row < len(grid) && col >= 0 && col < len(grid[row])
}

func Abs(n int64) int64 {
	if n < 0 {
		return -n
	}
	return n
}

func Assert(condition bool, msg string) {
	if !condition {
		panic(msg)
	}
}
