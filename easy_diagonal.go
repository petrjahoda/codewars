package main

import "math/big"

func CodeWarsEasyDiagonal() {
	println(diagonal(7, 0)) // => 8
	println(diagonal(7, 1)) // => 28
	println(diagonal(20, 3)) // => 5985
	println(diagonal(20, 4)) // => 20349
}

func diagonal(maxRows int, targetColumn int) int {
	//sum := 0
	//for row := 0; row <= maxRows; row++ {
	//	temp := 1
	//	for column := 0; column <= row; column++ {
	//		if column == 0 || row == 0 {
	//			temp = 1
	//		} else {
	//			temp = temp * (row - column + 1) / column
	//		}
	//		if column == targetColumn {
	//			sum += temp
	//		}
	//	}
	//}
	//return sum
	var z big.Int
	return int(z.Binomial(int64(maxRows+1), int64(targetColumn+1)).Int64())
}
