package main

import (
	"strconv"
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
	var sum uint64
	var temp uint64
	for row := 0; row <= n; row++ {
		sum = 0
		temp = 1
		for column := 0; column <= row; column++ {
			if column == 0 || row == 0 {
				temp = 1
			} else {
				temp = temp * (uint64(row) - uint64(column) + 1) / uint64(column)
			}
			sum += temp * temp
		}
	}
	return strconv.FormatUint(sum, 10)
	//todo: problem with integer overflow, convert to bigint
}
