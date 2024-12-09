package day8

import (
	"adventOfCode/util"
	"fmt"
)

func calculateAntinodes2(gridSize util.Vec2, pointA util.Vec2, pointB util.Vec2) (result []util.Vec2) {
	diff := util.SubVec2(pointA, pointB)

	currentPoint := pointA
	for util.WithinGrid(util.Vec2{}, gridSize, currentPoint) {
		result = append(result, currentPoint)
		currentPoint = util.AddVec2(currentPoint, diff)
	}

	currentPoint = pointB
	for util.WithinGrid(util.Vec2{}, gridSize, currentPoint) {
		result = append(result, currentPoint)
		currentPoint = util.SubVec2(currentPoint, diff)
	}

	return
}

func Task2() {
	data, gridSize := loadAndParseData()

	antiNodes := map[util.Vec2]bool{}

	for _, antennas := range data {
		for idx, antennaA := range antennas {
			for _, antennaB := range antennas[idx+1:] {
				for _, point := range calculateAntinodes2(gridSize, antennaA, antennaB) {
					antiNodes[point] = true
				}
			}
		}
	}
	fmt.Println(len(antiNodes))
}
