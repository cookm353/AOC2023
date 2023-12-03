package main

import (
	"fmt"

	day1 "github.com/cookm353/AdventOfCode/2023/Day1"
	day2 "github.com/cookm353/AdventOfCode/2023/Day2"
)

func main() {

	for i := 0; i < 10; i++ {
		// sumCalibrationValues := day1.FindCalibrationValue()
		sumCalibrationValues := day1.Part1()
		fmt.Println("Sum of calibration values:", sumCalibrationValues)
	}

	// game := "Game 23: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue"

	// re := regexp.MustCompile(`[0-9]+:`)
	// gameNumber := string(re.Find([]byte(game)))
	// fmt.Println(gameNumber[:len(gameNumber)-1])
	fmt.Println()
	day2.Part1()
}
