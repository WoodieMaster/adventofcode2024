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
