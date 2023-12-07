package day3

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// var charMap = map[int]int{}

func Part1() {
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

	input := strings.Split(testInput, "\n")
	partNumbers := []int{}

	for i, line := range input {
		isPartNumber := false
		currNumber := ""
		fmt.Println(strings.TrimSpace(line))
		for j, char := range line {

			// Build string for current number
			if unicode.IsNumber(char) {
				currNumber += string(char)

				if !isPartNumber {
					isPartNumber = checkIfValid(i, j, input)
				}
			}
			// Check if complete number is a part number
			if currNumber != "" && (!unicode.IsNumber(char) || j == len(line)-1) {
				partNumber, _ := strconv.Atoi(currNumber)
				currNumber = ""
				fmt.Println(partNumber, isPartNumber)

				if isPartNumber {
					partNumbers = append(partNumbers, partNumber)
				}
				isPartNumber = false
			}
		}
		fmt.Println()
	}

	var sum int
	fmt.Println(partNumbers)
	for _, partNumber := range partNumbers {
		sum += partNumber
	}
	fmt.Println(sum)
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
