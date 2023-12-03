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
	input := readInput("Day2/input1.txt")

	reBlue := regexp.MustCompile(`[0-9]+.blue`)
	reRed := regexp.MustCompile(`[0-9]+.red`)
	reGreen := regexp.MustCompile(`[0-9]+.green`)

	for idx, line := range input {
		numRedCubes := breakUpGames(reRed.FindAll(line, -1))
		numBlueCubes := breakUpGames(reBlue.FindAll(line, -1))
		numGreenCubes := breakUpGames(reGreen.FindAll(line, -1))

		if numRedCubes <= 12 && numBlueCubes <= 14 && numGreenCubes <= 13 {
			sumValidGameID += idx + 1
		}
	}
	fmt.Println("Sum of game IDs;", sumValidGameID)
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

/* Return number of cubes played for a given color */
func breakUpGames(games [][]byte) int {
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
