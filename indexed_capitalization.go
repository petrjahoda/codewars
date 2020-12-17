package main

import (
	"sort"
	"strings"
)

//Given a string and an array of integers representing indices, capitalize all letters at the given indices.
//For example:
//
//capitalize("abcdef",[1,2,5]) = "aBCdeF"
//capitalize("abcdef",[1,2,5,100]) = "aBCdeF". There is no index 100.
//The input will be a lowercase string with no spaces and an array of digits.

func CodewarsIndexedCapitalization() {
	Capitalize("abcdef", []int {1,2,5,100})
}

func Capitalize(inputString string, inputArray []int) string {
	sort.Ints(inputArray)
	var outputString string
	for position, character := range inputString {
		positionIndex := sort.SearchInts(inputArray, position)
		if positionIndex < len(inputArray) && inputArray[positionIndex] == position {
			outputString += strings.ToUpper(string(character))
		} else {
			outputString += string(character)
		}
	}
	return outputString
}
