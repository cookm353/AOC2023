package day3

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

// var charMap = map[int]int{}

func Part1() {
	/*
		testInput := `467..114..
			...*......
			..35..633.
			......#...
			617*......
			.....+.58.
			..592.....
			......755.
			...$.*....
			.664.598..`

		inputStr := strings.Split(testInput, "\n")
	*/
	input := readInput("Day3/input.txt")
	partNumbers := []int{}

	for i, line := range input {
		// symbols := []string{}
		fmt.Println(string(line))
		lineSymbols := ""
		for j, char := range line {
			adjacentNumbers := []int{}
			currChar := rune(char)
			if (unicode.IsPunct(currChar) && string(currChar) != ".") || unicode.IsSymbol(currChar) {
				fmt.Println(string(currChar))
				adjacentNumbers = findNumbers(i, j, input)
				lineSymbols += string(currChar)
			}

			for _, num := range adjacentNumbers {
				fmt.Println("Adjacent numbers:", num)
				partNumbers = append(partNumbers, num)
			}
		}
		lineSymbols = ""
	}

	var sum int
	for partNumber := range partNumbers {
		sum += partNumber
	}
	fmt.Println("Sum", sum)
}

func readInput(path string) [][]byte {
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	input := [][]byte{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := []byte(scanner.Text())
		input = append(input, line)
	}

	return input
}

/* Find numbers adjacent to character */
func findNumbers(i, j int, input [][]byte) []int {
	partNums := []int{}
	iMax := len(input) - 1
	jMax := len(input[0]) - 1

	iVals := []int{i}
	jVals := []int{j}

	// Build arrays of valid adjacent indices
	if i > 0 {
		iVals = append(iVals, i-1)
	}
	if i < iMax {
		iVals = append(iVals, i+1)
	}
	if j > 0 {
		jVals = append(jVals, j-1)
	}
	if j < jMax {
		jVals = append(jVals, j+1)
	}

	// Loop over all adjacent elements, find any numbers, and add to slice
	for _, iIdx := range iVals {
		for _, jIdx := range jVals {
			char := input[iIdx][jIdx]
			if unicode.IsNumber(rune(char)) {
				left := jIdx
				right := jIdx

				if j == jIdx {
					break
				}

				for left > 0 && unicode.IsNumber(rune(input[iIdx][left-1])) {
					left--
				}

				for right < jMax && unicode.IsNumber(rune(input[iIdx][right+1])) {
					right++
				}

				partNum, _ := strconv.Atoi(string(input[iIdx][left : right+1]))
				partNums = append(partNums, partNum)
			}
		}
	}

	return partNums
}

/* Check surrounding  */
func checkIfValid(i, j int, input []string) bool {
	iMax := len(input) - 1
	jMax := len(input[0]) - 1

	iVals := []int{i}
	jVals := []int{j}

	// Build arrays of valid adjacent indices
	if i > 0 {
		iVals = append(iVals, i-1)
	}
	if i < iMax {
		iVals = append(iVals, i+1)
	}
	if j > 0 {
		jVals = append(jVals, j-1)
	}
	if j < jMax {
		jVals = append(jVals, j+1)
	}

	// Loop over adjacent elements and check for characters
	for _, iIdx := range iVals {
		for _, jIdx := range jVals {
			char := input[iIdx][jIdx]
			if !unicode.IsNumber(rune(char)) && string(char) != "." {
				fmt.Println(string(char), i, j, "foo")
				return true
			}
		}
	}

	return false
}
