package main

func CodeWarsConsecutiveFibNumbers() {

	productFib(714) // should return [21, 34, 1],
	productFib(800) // should return [34, 55, 0],
}

func productFib(prod uint64) [3]uint64 {
	var previousFib uint64 = 1
	var actualFib uint64 = 1
	var match uint64 = 0
	for i := 0; i < int(prod); i++ {
		if actualFib*previousFib == prod {
			match = 1
			break
		}
		if actualFib*previousFib > prod {
			break
		}
		temp := actualFib
		actualFib = actualFib + previousFib
		previousFib = temp
	}
	return [3]uint64{previousFib, actualFib, match}
}
