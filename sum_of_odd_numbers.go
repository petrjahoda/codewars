package main

import "time"

//Given the triangle of consecutive odd numbers:
//
//1: 1
//2: 3     5
//3: 7     9    11
//4: 13    15    17    19
//5: 21    23    25    27    29
//6: 31
//...
//Calculate the row sums of this triangle from the row index (starting at index 1) e.g.:
//
//rowSumOddNumbers(1); // 1
//rowSumOddNumbers(2); // 3 + 5 = 8

func CodewarsSumOfOddNumbers() {
	timerEasy := time.Now()
	println(RowSumOddNumbersEasy(90))
	println("Takes " + time.Since(timerEasy).String())
	timer := time.Now()
	println(RowSumOddNumbers(90))
	println("Takes " + time.Since(timer).String())
}

func RowSumOddNumbersEasy(n int) int {
	return n * n * n
}

func RowSumOddNumbers(n int) int {
	result := 0
	firstNumber := 1 + (n-1)*n
	for i := 0; i < n; i++ {
		result += firstNumber
		firstNumber += 2
	}
	return result
}
