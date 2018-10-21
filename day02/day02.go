package day02

import (
	"fmt"
	"strconv"
	"strings"
)

func Solve02(input []byte)  {
	checksum := 0
	divsum := 0
	for _, m := range strings.Split(string(input), "\n") {
		max, min := highestAndLowest(m)
		checksum += (max - min)
		divsum += evenlyDivisible(m)
	}
	fmt.Printf("Check sum is: %v \n", checksum)
	fmt.Printf("Div sum is: %v \n", divsum)

}

func highestAndLowest(input string) (int, int)  {
	line := strings.Split(input, "\t")
	var max int
	var min int
	for _, value := range line  {
		actual, err := strconv.Atoi(value)

		if err != nil {
			panic(err)
		}

		if max == 0 {
			max = actual
		}

		if min == 0 {
			min = actual
		}

		if actual < min {
			min = actual
		}

		if actual > max {
			max = actual
		}
	}

	return max, min

}

func evenlyDivisible(input string) (int) {
	line := strings.Split(input, "\t")

	for _, value := range line  {
		actual, err := strconv.Atoi(value)

		if err != nil {
			panic(err)
		}

		for _, number := range line  {
			actualNumber, _ := strconv.Atoi(number)
			if actual == actualNumber {
				continue
			}

			if actual % actualNumber == 0 {
				if actual > actualNumber{
					return actual / actualNumber
				}
				return actualNumber / actual
			}
		}
	}
	return 0
}
