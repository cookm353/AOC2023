package main

import (
	"fmt"

	day1 "github.com/cookm353/AdventOfCode/2023/Day1"
	day2 "github.com/cookm353/AdventOfCode/2023/Day2"
	day3 "github.com/cookm353/AdventOfCode/2023/Day3"
)

type foo interface {
	int | float64
}

func add[T foo](a, b T) T {
	return a + b
}

func main() {
	fmt.Println("Day 1")
	fmt.Println("-----")
	day1.Part1()
	fmt.Println("\nDay 2")
	fmt.Println("-----")
	day2.Part1()
	fmt.Println("\nDay 3")
	fmt.Println("-----")
	day3.Part1()

	fmt.Println(add(3.141, 69))
}
