package day04

import (
	"fmt"
	"reflect"
	"strings"
)

func Solve04(input []byte)  {
	solutionPart1, solutionPart2 := part1(input)
	fmt.Printf("Solution for first part: %v\n", solutionPart1)
	fmt.Printf("Solution for first part: %v\n", solutionPart2)
}

func part1(input []byte) (int, int) {
	var validPassPhrases = 0
	var validNoAnagrammaPassPhrases = 0
	for _, m := range strings.Split(string(input), "\n") {
		var isValid bool = true
		var isValidNoAnagramma bool = true
		for k1, word := range strings.Split(m, " ") {
			for k2, word2 := range strings.Split(m, " ") {
				if word == word2 && k1 != k2 {
					isValid = false
					isValidNoAnagramma = false
					break
				}
				if isAnagram(word, word2) && k1 != k2 {
					isValidNoAnagramma = false
					break
				}
			}
		}
		if isValid {
			validPassPhrases++
		}
		if isValidNoAnagramma {
			validNoAnagrammaPassPhrases++
		}
	}

	return validPassPhrases, validNoAnagrammaPassPhrases
}

func isAnagram(word, word2 string) bool {
	word1Parts := map[string]int{}
	word2Parts := map[string]int{}
	for _, v := range word {
		if word1Parts[string(v)] == 1 {
			word1Parts[string(v)]++
		} else {
			word1Parts[string(v)] = 1
		}
	}
	for _, v := range word2 {
		if word2Parts[string(v)] == 1 {
			word2Parts[string(v)]++
		} else {
			word2Parts[string(v)] = 1
		}
	}
	return reflect.DeepEqual(word1Parts, word2Parts)
	/*for k, _ := range word1Parts {
		if word1Parts[k] != word2Parts[k] {
			return false
		}
	}*/

	//return true
}

