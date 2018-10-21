package day03

import (
	"fmt"
	"math"
	"strconv"
)

type coordinate struct {
	number int
	xCord int
	yCord int
}

func Solve03(input []byte) {
	inputNum, err := strconv.Atoi(string(input))
	if err != nil {
		panic(err)
	}

	solutionPart1 := part1(inputNum)
	solutionPart2 := part2(inputNum)
	fmt.Printf("Solution for first part: %v\n", solutionPart1)
	fmt.Printf("Solution for first part: %v\n", solutionPart2)
}

func part1(inputNum int) int  {
	width := int(math.Ceil(math.Sqrt(float64(inputNum))))
	if width%2 == 0 {
		width++
	}
	maxAmount := width * width

	minRow, minCol := 0, 0
	maxCol := width - 1
	maxRow := width - 1
	currentCol := width - 1
	currentRow := width - 1
	direction := "left"

	for maxAmount != inputNum {
		switch direction {
		case "left":
			if currentCol > minCol {
				currentCol--
				maxAmount--
			} else {
				direction = "up"
				minCol++
			}
		case "up":
			if currentRow > minRow {
				currentRow--
				maxAmount--
			} else {
				direction = "right"
				minRow++
			}
		case "right":
			if currentCol < maxCol {
				currentCol++
				maxAmount--
			} else {
				direction = "down"
				maxCol--
			}
		case "down":
			if currentRow < maxRow {
				currentRow++
				maxAmount--
			} else {
				direction = "left"
				maxRow--
			}
		}
	}
	a := int(math.Abs(float64(currentRow-width/2)))
	b := int(math.Abs(float64(currentCol-width/2)))

	return a + b
}

var mapping []coordinate
var currentNum = 0

func part2(inputNum int) int {
	currentMaxRow := 1
	currentMaxCol := 1
	currentMinRow := -1
	currentMinCol := -1
	currentX := 0
	currentY := 0
	direction := "right"

	for currentNum < inputNum {
		if currentNum == 0 {
			mapping = append(mapping, coordinate{
				number: 1,
				xCord: 0,
				yCord: 0,
			})
			currentNum++
		}

		switch direction {
		case "right":
			if currentX < currentMaxCol {
				currentX++
				extendMapping(currentX, currentY)
				currentNum = calculateCurrentValue()
				mapping[len(mapping) - 1].number = currentNum
			} else {
				direction = "up"
				currentMaxCol++
			}
			break
		case "up":
			if currentY < currentMaxRow{
				currentY++
				extendMapping(currentX, currentY)
				currentNum = calculateCurrentValue()
				mapping[len(mapping) - 1].number = currentNum
			} else {
				direction = "left"
				currentMaxRow++
			}
			break
		case "left":
			if currentX > currentMinCol {
				currentX--
				extendMapping(currentX, currentY)
				currentNum = calculateCurrentValue()
				mapping[len(mapping) - 1].number = currentNum
			} else {
				direction = "down"
				currentMinCol--
			}
		case "down":
			if currentY > currentMinRow {
				currentY--
				extendMapping(currentX, currentY)
				currentNum = calculateCurrentValue()
				mapping[len(mapping) - 1].number = currentNum
			} else {
				direction = "right"
				currentMinRow--
			}
		}
	}

	return currentNum
}

func extendMapping(currentX, currentY int) {
	mapping = append(mapping, coordinate{
		number: 0,
		xCord: currentX,
		yCord: currentY,
	})
}

func calculateCurrentValue() int  {
	var coordinatesAroundCurrent []coordinate
	var currentNumber int
	lastCoordinate := mapping[len(mapping) - 1]

	offset := []int{-1, 0, 1}
	for _, v := range offset{
		for _, v2 := range offset{
			xCord := lastCoordinate.xCord + v
			yCord := lastCoordinate.yCord + v2
			coordinatesAroundCurrent = append(coordinatesAroundCurrent, coordinate{
				number: 0,
				xCord: xCord,
				yCord: yCord,
			})
		}
	}

	for _, v := range coordinatesAroundCurrent {
		for _, mapValue := range mapping {
			if mapValue.xCord == v.xCord && mapValue.yCord == v.yCord {
				currentNumber += mapValue.number
			}
		}
	}

	return currentNumber
}
