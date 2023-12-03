package main

import (
	"fmt"
	"strings"

	day1 "github.com/cookm353/AdventOfCode/2023/Day1"
)

func main() {

	for i := 0; i < 10; i++ {
		// sumCalibrationValues := day1.FindCalibrationValue()
		sumCalibrationValues := day1.Part1()
		fmt.Println(sumCalibrationValues)
	}

	line := "eightwoneight"

	for k, v := range day1.DigitMap {
		if strings.Contains(line, k) {
			line = strings.ReplaceAll(line, k, v)
		}
	}
	fmt.Println(line)
}
