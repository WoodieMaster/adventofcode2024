package day7

import (
	"fmt"
)

func Task1() {
	data := loadAndParseData()
	result := 0

	for _, operation := range data {
		if hasValidCalculation(operation, Plus, Mult) {
			result += operation.Result
		}
	}

	fmt.Println(result)
}
