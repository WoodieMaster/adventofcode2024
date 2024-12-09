package day7

import (
	"fmt"
)

func Task2() {
	data := loadAndParseData()
	result := 0

	for _, operation := range data {
		if hasValidCalculation(operation, Plus, Mult, Concat) {
			result += operation.Result
		}
	}

	fmt.Println(result)
}
