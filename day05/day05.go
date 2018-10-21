package day05

import (
	"fmt"
	"strconv"
	"strings"
)

func Solve05(input []byte)  {
	solutionPart1 := part1(input)
	solutionPart2 := part2(input)
	fmt.Printf("Solution for first part: %v\n", solutionPart1)
	fmt.Printf("Solution for first part: %v\n", solutionPart2)
}

func part1(input []byte) int {
	var steps []int

	list := strings.Split(string(input), "\n")
	for _, v := range list {
		value, _ := strconv.Atoi(v)
		steps = append(steps, value)
	}

	currentIndex := 0
	lastKey := len(steps) - 1
	counter := 0

	for currentIndex <= lastKey {
		move := steps[currentIndex]
		steps[currentIndex]++
		currentIndex += move
		counter++
	}


	return counter
}

func part2(input []byte) int {
	var steps []int

	list := strings.Split(string(input), "\n")
	for _, v := range list {
		value, _ := strconv.Atoi(v)
		steps = append(steps, value)
	}

	currentIndex := 0
	lastKey := len(steps) - 1
	counter := 0

	for currentIndex <= lastKey {
		move := steps[currentIndex]
		if move >= 3 {
			steps[currentIndex]--
		} else {
			steps[currentIndex]++
		}
		currentIndex += move
		counter++
	}


	return counter
}