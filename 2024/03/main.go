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

	fmt.Println("first star:", firstStar(string(input)))
	fmt.Println("second star:", secondStar(string(input)))
}

func firstStar(input string) int {
	sum := 0

	parseMul(input, &sum)

	return sum
}

func parseMul(s string, sum *int) {
	matches := strings.Split(s, "mul")
	for i := 1; i < len(matches); i++ { // Start at 1 to skip first split
		// fmt.Println(matches[i])
		if len(matches[i]) < 5 { // Need at least (d,d)
			continue
		}
		if matches[i][0] != '(' {
			continue
		}
		closeParen := strings.Index(matches[i], ")")
		if closeParen == -1 {
			continue
		}
		nums := strings.Split(matches[i][1:closeParen], ",")
		if len(nums) != 2 {
			continue
		}
		if len(nums[0]) < 1 || len(nums[0]) > 3 || len(nums[1]) < 1 || len(nums[1]) > 3 {
			continue
		}
		n1, err1 := strconv.Atoi(nums[0])
		n2, err2 := strconv.Atoi(nums[1])
		if err1 == nil && err2 == nil {
			*sum += n1 * n2
		}
	}
}
