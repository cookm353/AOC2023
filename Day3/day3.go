package day3

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

var charMap = map[int]int{}

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

	test := strings.Split(testInput, "\n")

	for i, line := range test {
		for _, char := range line {
			fmt.Println(unicode.IsSymbol(rune(char)), string(char))
		}
		fmt.Println(i, strings.TrimSpace(line))
	}
	// fmt.Println(testInput)
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

/*
Read input
Loop over each line
	for i, line := range lines
Find indices of each symbol
		for j, char := range line
			if char is a symbol


Check if there are numbers

hash map
	key holds i
	val holds array of j
*/
