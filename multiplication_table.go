package main

import "fmt"

//Your task, is to create NxN multiplication table, of size provided in parameter.
//for example, when given size is 3:
//1 2 3
//2 4 6
//3 6 9
//for given example, the return value should be: [[1,2,3],[2,4,6],[3,6,9]]


func CodewarsMultiplicationTable()  {
	result := MultiplicationTable(6)
	fmt.Println(result)
}


func MultiplicationTable(size int) [][]int {
	var first [][]int
	for i := 0; i < size; i++ {
		var second []int
		for j := 0; j < size; j++ {
			second = append(second, (i+1)*(j+1))
		}
		first = append(first, second)
	}
	return first
}