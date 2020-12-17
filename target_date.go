package main

import (
	"time"
)

func CodeWarsTargetDate() {
	println(DateNbDays(100, 101, 0.98)) // "2017-01-01" (366 days)
	println(DateNbDays(100, 150, 2.00))    // "2035-12-26" (7299 days)
}

func DateNbDays(a0 float32, a float32, p float32) string {
	startDate := time.Date(2016, 01, 01, 0, 0, 0, 0, time.UTC)
	elapsedDays := 0
	for a0 < a {
		a0 += a0 * (p/36000)
		elapsedDays++
	}
	return startDate.AddDate(0, 0, elapsedDays).Format("2006-01-01")
}
