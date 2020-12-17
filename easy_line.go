package main

import (
	"math/big"
)

func CodeWarsEasyLine() {
	println(easyline(0))   //=> 1
	println(easyline(1))   //=> 2
	println(easyline(4))   //=> 70
	println(easyline(7))   //=> 3432
	println(easyline(13))  //=> 1040060
	println(easyline(50))  //=> 100891344545564193334812497256
	println(easyline(100)) //=> 90548514656103281165404177077484163874504589675413336841320
}

func easyline(n int) string {
	sum := big.NewInt(0)
	temp := big.NewInt(0)
	result := big.NewInt(0)
	for row := 0; row <= n; row++ {
		result.SetUint64(0)
		sum.SetUint64(0)
		temp.SetUint64(1)
		for column := 0; column <= row; column++ {
			if column == 0 || row == 0 {
				temp.SetUint64(1)
			} else {
				integer := row - column + 1
				multiplicator := big.NewInt(int64(integer))
				temp.Mul(temp, multiplicator)
				columnIndicator := big.NewInt(int64(column))
				temp.Div(temp, columnIndicator)
			}
			result.Mul(temp, temp)
			sum.Add(sum, result)
		}
	}
	return sum.Text(10)
}
