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
	input := []string{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"}
	// input := readInput("Day2/input1.txt")

	reBlue := regexp.MustCompile(`[0-9]+.blue`)
	reRed := regexp.MustCompile(`[0-9]+.red`)
	reGreen := regexp.MustCompile(`[0-9]+.green`)

	for idx, line := range input {
		game := []byte(line)

		numRedCubes := breakUpGames(reRed.FindAll(game, -1))
		numBlueCubes := breakUpGames(reBlue.FindAll(game, -1))
		numGreenCubes := breakUpGames(reGreen.FindAll(game, -1))
		// numRedCubes := breakUpGames(reRed.FindAll(line, -1))
		// numBlueCubes := breakUpGames(reBlue.FindAll(line, -1))
		// numGreenCubes := breakUpGames(reGreen.FindAll(line, -1))

		if numRedCubes <= 12 && numBlueCubes <= 14 && numGreenCubes <= 13 {
			sumValidGameID += idx + 1
		}
	}
	fmt.Println(sumValidGameID)
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
	cubeCount := 0

	for _, game := range games {
		numString := string(reNumGames.Find(game))
		numCubes, _ := strconv.Atoi(numString)
		cubeCount += numCubes
	}

	return cubeCount
}
