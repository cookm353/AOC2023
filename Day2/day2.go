package day2

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func Part1() {
	sumValidGameID := 0
	sumSetPowers := 0
	input := readInput("Day2/input.txt")

	reBlue := regexp.MustCompile(`[0-9]+.blue`)
	reRed := regexp.MustCompile(`[0-9]+.red`)
	reGreen := regexp.MustCompile(`[0-9]+.green`)

	for idx, line := range input {
		maxRedCubes := findMaxCubes(reRed.FindAll(line, -1))
		maxBlueCubes := findMaxCubes(reBlue.FindAll(line, -1))
		maxGreenCubes := findMaxCubes(reGreen.FindAll(line, -1))

		if maxRedCubes <= 12 && maxBlueCubes <= 14 && maxGreenCubes <= 13 {
			sumValidGameID += idx + 1
		}

		sumSetPowers += (maxRedCubes * maxBlueCubes * maxGreenCubes)
	}
	fmt.Println("Sum of game IDs;", sumValidGameID)
	fmt.Println("Sum of set powers: ", sumSetPowers)
}

/* Read input from file and split into 2-D byte array */
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

/* Find greatest cubes played for a given color */
func findMaxCubes(games [][]byte) int {
	reNumGames := regexp.MustCompile(`[0-9]+`)
	maxCubes := 0

	for _, game := range games {
		numString := string(reNumGames.Find(game))
		numCubes, _ := strconv.Atoi(numString)
		if numCubes > maxCubes {
			maxCubes = numCubes
		}
	}

	return maxCubes
}
