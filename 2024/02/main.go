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
