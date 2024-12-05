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
	numbers, count := createRules(lines)

	fmt.Println("first star:", firstStar(lines, numbers, count))
}

func createRules(lines []string) (map[int][]int, int) {
	count := 0
	numbers := make(map[int][]int)
	for _, line := range lines {
		if len(line) < 1 {
			break
		}
		nums := strings.Split(line, "|")
		if len(nums) != 2 {
			continue
		}
		num1, err1 := strconv.Atoi(nums[0])
		num2, err2 := strconv.Atoi(nums[1])
		if err1 != nil || err2 != nil {
			continue
		}
		numbers[num1] = append(numbers[num1], num2)
		count++
	}

	return numbers, count
}

func contains(arr []int, value int) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}
	return false
}

func firstStar(lines []string, numbers map[int][]int, count int) int {
	sum := 0
	for i := count + 1; i < len(lines); i++ {
		nums := strings.Split(lines[i], ",")
		found := false
		appArr := []int{}
		num0, err := strconv.Atoi(nums[0])
		if err != nil {
			continue
		}
		appArr = append(appArr, num0)
		for c := 1; c < len(nums); c++ {
			num, err := strconv.Atoi(nums[c])
			if err != nil {
				continue
			}
			for _, n := range numbers[num] {
				if contains(appArr, n) {
					found = true
					break
				}
			}
			if found {
				break
			}
			appArr = append(appArr, num)
		}
		if !found {
			sum += appArr[int(len(appArr))/2]
		}
	}
	return sum
}
