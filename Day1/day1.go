package day1

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var DigitMaps = map[string]string{
	"zero":  "0",
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

type num struct {
	idx    int
	number string
}

func FindCalibrationValue() int {
	fileName := "Day1/input.txt"

	data, err := os.ReadFile(fileName)

	if err != nil {
		panic(err)
	}

	nextNumber := ""
	calibrationValues := []int{}
	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		firstNum := num{len(line), ""}
		lastNum := num{0, ""}
		fmt.Print(line + " = ")

		// Loop over each number
		for text, digit := range DigitMaps {
			if strings.Index(line, text) < firstNum.idx && strings.Contains(line, text) {
				firstNum.idx = strings.Index(line, text)
				firstNum.number = digit
			} else if strings.Index(line, digit) < firstNum.idx && strings.Contains(line, digit) {
				firstNum.idx = strings.Index(line, digit)
				firstNum.number = digit
			}
			if strings.LastIndex(line, text) > lastNum.idx && strings.Contains(line, text) {
				lastNum.idx = strings.LastIndex(line, text)
				lastNum.number = digit
			} else if strings.LastIndex(line, digit) > lastNum.idx && strings.Contains(line, digit) {
				lastNum.idx = strings.LastIndex(line, digit)
				lastNum.number = digit
			}
		}
		fmt.Printf("%v%v\n", firstNum.number, lastNum.number)

		nextNumber = firstNum.number + lastNum.number
		fmt.Printf("%v\t%v\t%v\n", firstNum.number, lastNum.number, nextNumber)

		value, _ := strconv.Atoi(nextNumber)
		calibrationValues = append(calibrationValues, value)
		nextNumber = ""
	}

	/*
		loop over every character in input
		add numeric characters to temp string
		when we encounter a line break or EOF...
			if it has more than 2 digits...
				create new string made of first and last digits
			convert to an int, add to list of calibration values
	*/

	// for idx, char := range data {
	// 	if unicode.IsNumber(rune(char)) {
	// 		nextNumber += string(char)
	// 	} else if char == 10 || idx == len(data)-1 {
	// 		if len(nextNumber) > 2 {
	// 			nextNumber = nextNumber[0:1] + nextNumber[len(nextNumber)-1:]
	// 		}
	// 		value, _ := strconv.Atoi(nextNumber)
	// 		calibrationValues = append(calibrationValues, value)
	// 		nextNumber = ""
	// 	}
	// }

	sumCalibrationValues := 0

	for _, num := range calibrationValues {
		fmt.Printf("%v\t+\t%v\t=\t%v\n", num, sumCalibrationValues, num+sumCalibrationValues)
		sumCalibrationValues += num
	}
	fmt.Println(sumCalibrationValues)

	// line := "3nine4fourjclspd152rpv"
	// firstNum := num{len(line), ""}
	// lastNum := num{0, ""}
	// for text, digit := range DigitMaps {
	// 	if strings.Index(line, text) < firstNum.idx && strings.Contains(line, text) {
	// 		firstNum.idx = strings.Index(line, text)
	// 		firstNum.number = digit
	// 	}
	// 	if strings.Index(line, digit) < firstNum.idx && strings.Contains(line, digit) {
	// 		firstNum.idx = strings.Index(line, digit)
	// 		firstNum.number = digit
	// 	}
	// 	if strings.LastIndex(line, text) > lastNum.idx && strings.Contains(line, text) {
	// 		lastNum.idx = strings.LastIndex(line, text)
	// 		lastNum.number = digit
	// 	}
	// 	if strings.LastIndex(line, digit) > lastNum.idx && strings.Contains(line, digit) {
	// 		lastNum.idx = strings.LastIndex(line, digit)
	// 		lastNum.number = digit
	// 	}
	// }
	// fmt.Printf("%v\t%v\n", firstNum.number, lastNum.number)

	return sumCalibrationValues
}

func PullNumbers(text []string) {
	for line := range text {
		fmt.Println(line)
	}
}
