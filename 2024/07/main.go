package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// input, err := os.ReadFile("test.txt")
	input, err := os.ReadFile("puzzle.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(input), "\n")

	fmt.Println("first star:", firstStar(lines))
	// fmt.Println("second star:", secondStar(startMap, startPos))
}

func firstStar(lines []string) int {
	sum := 0
	for _, line := range lines {
		parts := strings.Split(line, ": ")
		result, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		numbers := strings.Split(parts[1], " ")
		numbersInt := make([]int, len(numbers))
		for i, number := range numbers {
			numbersInt[i], err = strconv.Atoi(number)
			if err != nil {
				panic(err)
			}
		}
		isValid := tryMultiply(result, numbersInt)
		if isValid {
			sum += result
		}
	}
	return sum
}

func tryMultiply(target int, numbers []int) bool {
	if len(numbers) == 1 {
		return numbers[0] == target
	}

	isValid := tryMultiply(target, append([]int{numbers[0] + numbers[1]}, numbers[2:]...))
	if isValid {
		return isValid
	}
	isValid = tryMultiply(target, append([]int{numbers[0] * numbers[1]}, numbers[2:]...))
	if isValid {
		return isValid
	}
	concatenated := strconv.Itoa(numbers[0]) + strconv.Itoa(numbers[1])
	concatenatedInt, err := strconv.Atoi(concatenated)
	if err != nil {
		panic(err)
	}
	return tryMultiply(target, append([]int{concatenatedInt}, numbers[2:]...))
}
