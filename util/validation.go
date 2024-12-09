package util

func Must[T any](data T, err error) T {
	if err != nil {
		panic(err)
	}

	return data
}

func Assert(condition bool, msg string) {
	if !condition {
		panic(msg)
	}
}

func Permutations[T any](permutationSize int, values ...T) <-chan []T {
	Assert(len(values) >= 1, "values list must not be 0")

	c := make(chan []T)

	go func() {
		currentIndices := make([]int, permutationSize)
		currentPermutation := make([]T, permutationSize)

		for i := 0; i < permutationSize; i++ {
			currentIndices[i] = 0
			currentPermutation[i] = values[0]
		}

		for {
			c <- currentPermutation

			currentPermutation = make([]T, permutationSize)

			update := true

			for i := 0; i < permutationSize; i++ {
				if update {
					if currentIndices[i] < len(values)-1 {
						currentIndices[i]++
						update = false
					} else {
						currentIndices[i] = 0
					}
				}

				currentPermutation[i] = values[currentIndices[i]]
			}

			if update {
				close(c)
				return
			}
		}
	}()

	return c
}
