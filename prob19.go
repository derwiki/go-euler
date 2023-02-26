package main

import (
	"fmt"
	"time"
)

func main() {
	acc := 0
	startTime := time.Date(1901, time.January, 1, 0, 0, 0, 0, time.UTC)
	endTime := time.Date(2000, time.December, 31, 1, 0, 0, 0, time.UTC)
	for timeCursor := startTime; timeCursor.Before(endTime); timeCursor = timeCursor.AddDate(0, 0, 1) {
		if timeCursor.Weekday() == time.Weekday(0) && timeCursor.Day() == 1 {
			acc++
		}
	}
	fmt.Println(acc)
}
