package main

//You are given two arrays a1 and a2 of strings. Each string is composed with letters from a to z. Let x be any string in the first array and y be any string in the second array.
//
//Find max(abs(length(x) − length(y)))
//
//If a1 and/or a2 are empty return -1 in each language except in Haskell (F#) where you will return Nothing (None).
//
//Example:
//
//a1 = ["hoqq", "bbllkw", "oox", "ejjuyyy", "plmiis", "xxxzgpsssa", "xxwwkktt", "znnnnfqknaz", "qqquuhii", "dvvvwz"]
//a2 = ["cccooommaaqqoxii", "gggqaffhhh", "tttoowwwmmww"]
//mxdiflg(a1, a2) --> 13
//Bash note:
//
//input : 2 strings with substrings separated by ,
//output: number as a string

func CodeWarsMaximumLengthDifference() {
	a1 := []string{"hoqq", "bbllkw", "oox", "ejjuyyy", "plmiis", "xxxzgpsssa", "xxwwkktt", "znnnnfqknaz", "qqquuhii",
		"dvvvwz"}
	a2 := []string{"cccooommaaqqoxii", "gggqaffhhh", "tttoowwwmmww"}
	MxDifLg(a1, a2)
	a1 = []string{"ejjjjmmtthh", "zxxuueeg", "aanlljrrrxx", "dqqqaaabbb", "oocccffuucccjjjkkkjyyyeehh"}
	a2 = []string{"bbbaaayddqbbrrrv"}
	MxDifLg(a1, a2)
}

func MxDifLg(firstArray []string, secondArray []string) int {
	if len(firstArray) == 0 || len(secondArray) == 0 {
		return -1
	}
	firstStringLength := 0
	secondStringLength := 0
	resultLength := 0
	for _, firstString := range firstArray {
		for _, secondString := range secondArray {
			firstStringLength = len(firstString)
			secondStringLength = len(secondString)
			if secondStringLength-firstStringLength > resultLength {
				resultLength = secondStringLength - firstStringLength
			}
			if firstStringLength-secondStringLength > resultLength {
				resultLength = firstStringLength - secondStringLength
			}
		}
	}
	return resultLength
}
