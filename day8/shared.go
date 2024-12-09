package day8

import (
	"adventOfCode/util"
	"strings"
)

func validAntenna(id rune) bool {
	return (id >= '0' && id <= '9') ||
		(id >= 'a' && id <= 'z') ||
		(id >= 'A' && id <= 'Z')

}

func loadAndParseData() (result map[rune][]util.Vec2, gridSize util.Vec2) {
	text := util.LoadFile(8)
	lines := strings.Split(text, "\r\n")

	gridSize = util.Vec2{Row: len(lines), Col: len(lines[0])}

	result = map[rune][]util.Vec2{}

	for row, line := range lines {
		for col, char := range line {
			if validAntenna(char) {
				result[char] = append(result[char], util.Vec2{Row: row, Col: col})
			}
		}
	}

	return
}
