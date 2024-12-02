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

	numSafeFirst := 0
	numSafeSecond := 0

	for _, line := range lines {
		parts := strings.Split(line, " ")
		if len(parts) <= 1 {
			continue
		}

		if firstStar(parts) {
			numSafeFirst++
		}
		if secondStar(parts) {
			numSafeSecond++
		}
	}

	fmt.Println("first star:", numSafeFirst)
	fmt.Println("second star:", numSafeSecond)
}

func isSafe(diff int, isIncreasing bool) bool {
	if diff == 0 || diff > 3 || diff < -3 {
		return false
	} else if isIncreasing && diff < 0 {
		return false
	} else if !isIncreasing && diff > 0 {
		return false
	}
	return true
}

func firstStar(parts []string) bool {
	prev, _ := strconv.Atoi(parts[0])
	safe := true
	var isIncreasing bool
	for i := 1; i < len(parts); i++ {
		curr, _ := strconv.Atoi(parts[i])
		diff := curr - prev
		if i == 1 {
			isIncreasing = diff > 0
		}
		if !isSafe(diff, isIncreasing) {
			safe = false
			break
		}
		prev = curr
	}
	return safe
}

func secondStar(parts []string) bool {
	prev, _ := strconv.Atoi(parts[0])
	safe := true
	var isIncreasing bool
	for i := 1; i < len(parts); i++ {
		curr, _ := strconv.Atoi(parts[i])
		diff := curr - prev
		if i == 1 {
			isIncreasing = diff > 0
		}
		if !isSafe(diff, isIncreasing) {
			safe = false
			break
		}
		prev = curr
	}

	if !safe {
		for i := 0; i < len(parts); i++ {
			safe = secondStarTryAgain(parts, i)
			if safe {
				break
			}
		}
	}

	return safe
}

func secondStarTryAgain(parts []string, firstUnsafePos int) bool {
	var start int
	var prev int
	var isIncreasing bool
	if firstUnsafePos == 0 {
		prev, _ = strconv.Atoi(parts[1])
		start = 2
	} else {
		prev, _ = strconv.Atoi(parts[0])
		start = 1
	}
	if start == firstUnsafePos {
		start++
	}
	safe := true
	for i := start; i < len(parts); i++ {
		if i == firstUnsafePos {
			continue
		}
		curr, _ := strconv.Atoi(parts[i])
		diff := curr - prev
		if i == start {
			isIncreasing = diff > 0
		}
		if !isSafe(diff, isIncreasing) {
			safe = false
			break
		}
		prev = curr
	}
	return safe
}
