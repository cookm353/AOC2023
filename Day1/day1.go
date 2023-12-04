package day1

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var digitMap = map[string]string{
	"zero":  "z0o",
	"one":   "o1e",
	"two":   "t2o",
	"three": "t3e",
	"four":  "f4r",
	"five":  "f5e",
	"six":   "s6x",
	"seven": "s7n",
	"eight": "e8t",
	"nine":  "n9e",
}

func Part1() {
	fileName := "Day1/input.txt"
	lines := getLines(fileName)
	// someLines := lines[:10]

	calibrationValues := []int{}

	for _, line := range lines {
		// for _, line := range someLines {
		fmt.Println("original line:", line)
		line = strings.ToLower(line)
		line = convertNumbers(line)
		fmt.Println("converted numbers:", line)
		digits := getDigits(line)
		fmt.Println("digits:", digits)
		calibrationValues = append(calibrationValues, getCalibrationValue(digits))
		fmt.Println("calibration value:", getCalibrationValue(digits))
	}

	sumCalibrationValues := 0

	for _, calibrationValue := range calibrationValues {
		sumCalibrationValues += calibrationValue
	}

	fmt.Println("Sum of calibration values:", sumCalibrationValues)
}

/* Read text from input file */
func getLines(fileName string) []string {
	data, err := os.ReadFile(fileName)

	if err != nil {
		panic(err)
	}

	return strings.Split(string(data), "\n")
}

/* Convert spelled out number to digit (three to 3) */
func convertNumbers(line string) string {
	for k, v := range digitMap {
		if strings.Contains(line, k) {
			line = strings.ReplaceAll(line, k, v)
		}
	}

	return line
}

/* Extract the first and last digits from the string and cast to an int */
func getCalibrationValue(digits string) int {
	calibrationValueStr := string(digits[0]) + string(digits[len(digits)-1])
	calibrationValue, _ := strconv.Atoi(calibrationValueStr)

	return calibrationValue
}

/* Filter letters from line and return numbers */
func getDigits(line string) string {
	digits := ""
	for _, char := range line {
		if unicode.IsNumber(char) {
			digits += string(char)
		}
	}

	return digits
}

func FindCalibrationValue(lines []string) int {
	calibrationValues := []int{}
	for _, line := range lines {
		digits := []rune{}

		// Convert to digits
		for k, v := range digitMap {
			if strings.Contains(line, k) {
				line = strings.ReplaceAll(line, k, v)
			}
		}

		// Add numbers to slice
		for _, char := range line {
			if unicode.IsNumber(char) {
				digits = append(digits, char)
			}
		}

		if len(digits) == 1 {

		}
		nextNumber := fmt.Sprintf("%v%v", string(digits[0]), string(digits[len(digits)-1]))

		value, _ := strconv.Atoi(nextNumber)
		calibrationValues = append(calibrationValues, value)
		nextNumber = ""
	}

	sumCalibrationValues := 0

	for _, num := range calibrationValues {
		sumCalibrationValues += num
	}

	fmt.Println("Sum of calibration values:", sumCalibrationValues)

	return sumCalibrationValues
}
