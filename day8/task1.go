package day8

import (
	"adventOfCode/util"
	"fmt"
)

func calculateAntinodes1(gridSize util.Vec2, pointA util.Vec2, pointB util.Vec2) (result []util.Vec2) {
	diff := util.SubVec2(pointA, pointB)

	possiblePoint := util.AddVec2(pointA, diff)
	if util.WithinGrid(util.Vec2{}, gridSize, possiblePoint) {
		result = append(result, possiblePoint)
	}

	possiblePoint = util.SubVec2(pointB, diff)
	if util.WithinGrid(util.Vec2{}, gridSize, possiblePoint) {
		result = append(result, possiblePoint)
	}

	return
}

func Task1() {
	data, gridSize := loadAndParseData()

	antiNodes := map[util.Vec2]bool{}

	for _, antennas := range data {
		for idx, antennaA := range antennas {
			for _, antennaB := range antennas[idx+1:] {
				for _, point := range calculateAntinodes1(gridSize, antennaA, antennaB) {
					antiNodes[point] = true
				}
			}
		}
	}
	fmt.Println(len(antiNodes))
}
