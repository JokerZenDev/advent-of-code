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

	arr1 := []int{}
	arr2 := []int{}

	lines := strings.Split(string(input), "\n")

	for _, line := range lines {
		parts := strings.Split(line, "   ")
		if len(parts) != 2 {
			continue
		}
		num, _ := strconv.Atoi(parts[0])
		pos := findPos(arr1, num)
		if pos == -1 {
			arr1 = append(arr1, num)
		} else {
			arr1 = append(arr1[:pos], append([]int{num}, arr1[pos:]...)...)
		}
		num, _ = strconv.Atoi(parts[1])
		pos = findPos(arr2, num)
		if pos == -1 {
			arr2 = append(arr2, num)
		} else {
			arr2 = append(arr2[:pos], append([]int{num}, arr2[pos:]...)...)
		}
	}

	sum := 0
	for i := 0; i < len(arr1); i++ {
		if arr1[i] >= arr2[i] {
			sum += arr1[i] - arr2[i]
		} else {
			sum += arr2[i] - arr1[i]
		}
	}
	fmt.Println(sum)
}

func findPos(arr []int, num int) int {
	for i, v := range arr {
		if v >= num {
			return i
		}
	}
	return -1
}
