package main

import (
	"flag"
	"fmt"
	"github.com/ateee89/aof2017/day01"
	"github.com/ateee89/aof2017/day02"
	"github.com/ateee89/aof2017/day03"
	"github.com/ateee89/aof2017/day04"
	"github.com/ateee89/aof2017/day05"
	"io/ioutil"
)

var day int
var input string

func init() {
	flag.IntVar(&day, "day", 1, "Which day would you like to solve?")
	flag.StringVar(&input, "input", "none", "input for advent of code")
	flag.Parse()
}

func main() {
	err := validate()
	if err == 1 {
		return
	}

	dat, err1 := ioutil.ReadFile(input)
	check(err1)

	switch day {
	case 1:
		day01.Solve01(dat)
	case 2:
		 day02.Solve02(dat)
	case 3:
		 day03.Solve03(dat)
	case 4:
		 day04.Solve04(dat)
	case 5:
		 day05.Solve05(dat)
	default:
		fmt.Println("Not solved yet")
	}
}

func validate() int {
	if input == "none" {
		fmt.Println("No input were provided")
		return 1
	}
	return 0
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
