package day3

import (
	"bufio"
	"fmt"
	"os"
)

func Part1() {
	testInput := `
	467..114..
	...*......
	..35..633.
	......#...
	617*......
	.....+.58.
	..592.....
	......755.
	...$.*....
	.664.598..`

	fmt.Println(testInput)
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
