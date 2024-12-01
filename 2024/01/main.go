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

	firstStar(arr1, arr2)
	secondStar(arr1, arr2)
}

func firstStar(arr1, arr2 []int) {
	sum := 0
	for i := 0; i < len(arr1); i++ {
		if arr1[i] >= arr2[i] {
			sum += arr1[i] - arr2[i]
		} else {
			sum += arr2[i] - arr1[i]
		}
	}
	fmt.Println("first star:", sum)
}

func secondStar(arr1, arr2 []int) {
	sum := 0
	pos := 0
	for i := 0; i < len(arr1); i++ {
		finded, newPos := findEqualNumber(arr1, arr2[i], pos)
		pos = newPos
		sum += finded * arr2[i]
	}
	fmt.Println("second star:", sum)
}

func findPos(arr []int, num int) int {
	for i, v := range arr {
		if v >= num {
			return i
		}
	}
	return -1
}

func findEqualNumber(arr []int, num int, startPos int) (int, int) {
	finded := 0
	pos := startPos
	posIsChanged := false
	for i := startPos; i < len(arr); i++ {
		if arr[i] > num {
			break
		}
		if arr[i] == num {
			finded++
			if !posIsChanged {
				pos = i
				posIsChanged = true
			}
		}
	}
	return finded, pos
}
