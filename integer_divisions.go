package main

import "math"

func CodeWarsIntegersDivisions() {
	ListSquared(1, 250) // --> [[1, 1], [42, 2500], [246, 84100]]
	ListSquared(42, 1600) // --> [[42, 2500], [246, 84100]]
}

func ListSquared(m, n int) [][]int {
	result := make([][]int, 0)
	for i := m; i <= n; i++ {
		var goodIntegers []int
		for j := 1; j < n; j++ {
			if i%j == 0 {
				goodIntegers = append(goodIntegers, j*j)
			}
		}
		var sum int
		for _, integer := range goodIntegers {
			sum += integer
		}
		maxNumber := int(math.Sqrt(float64(sum)))
		if maxNumber*maxNumber == sum {
			if i <= sum {
				result = append(result, []int{i, sum})
			}
		}
	}
	return result
}
