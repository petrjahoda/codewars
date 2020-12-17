package main

func CodeWarsMaximumDivisors() {
	println(divisors(4))
	println(divisors(5))
	println(divisors(12))
	println(divisors(30))
	println(divisors(30000))
}

func divisors(ceiling int) int {
	sum := 1
	for i := 1; i <= ceiling/2; i++ {
		if ceiling % i == 0 {
			sum++
		}
	}
	return sum
}
