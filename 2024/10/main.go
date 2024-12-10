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

	hikingMap := createMap(lines)

	fmt.Println("first star:", firstStar(hikingMap))
	fmt.Println("second star:", secondStar(hikingMap))
}

func createMap(lines []string) [][]int {
	hikingMap := make([][]int, len(lines))
	for y, line := range lines {
		numbers := strings.Split(line, "")
		hikingMap[y] = make([]int, len(numbers))
		for x, number := range numbers {
			num, err := strconv.Atoi(number)
			if err != nil {
				panic(err)
			}
			hikingMap[y][x] = num
		}
	}
	return hikingMap
}

func printMap(hikingMap [][]int) {
	for _, line := range hikingMap {
		fmt.Println(line)
	}
}

func firstStar(hikingMap [][]int) int {
	sum := 0

	for y := 0; y < len(hikingMap); y++ {
		for x := 0; x < len(hikingMap[y]); x++ {
			if hikingMap[y][x] == 0 {
				endFound := startWalking(hikingMap, y, x)
				fmt.Println(y, x, endFound)
				sum += len(endFound)
			}
		}
	}

	return sum
}

func startWalking(hikingMap [][]int, y int, x int) [][2]int {
	if hikingMap[y][x] == 9 {
		return [][2]int{{y, x}}
	}

	endFound := [][2]int{}
	next := hikingMap[y][x] + 1

	if y > 0 && hikingMap[y-1][x] == next {
		// walk up
		toAdd := startWalking(hikingMap, y-1, x)
		for _, v := range toAdd {
			if !isIn(endFound, v[0], v[1]) {
				endFound = append(endFound, v)
			}
		}
	}
	if y < len(hikingMap)-1 && hikingMap[y+1][x] == next {
		// walk down
		toAdd := startWalking(hikingMap, y+1, x)
		for _, v := range toAdd {
			if !isIn(endFound, v[0], v[1]) {
				endFound = append(endFound, v)
			}
		}
	}
	if x > 0 && hikingMap[y][x-1] == next {
		// walk left
		toAdd := startWalking(hikingMap, y, x-1)
		for _, v := range toAdd {
			if !isIn(endFound, v[0], v[1]) {
				endFound = append(endFound, v)
			}
		}
	}
	if x < len(hikingMap[y])-1 && hikingMap[y][x+1] == next {
		// walk right
		toAdd := startWalking(hikingMap, y, x+1)
		for _, v := range toAdd {
			if !isIn(endFound, v[0], v[1]) {
				endFound = append(endFound, v)
			}
		}
	}
	return endFound
}

func isIn(endFound [][2]int, y int, x int) bool {
	for _, v := range endFound {
		if v[0] == y && v[1] == x {
			return true
		}
	}
	return false
}

func secondStar(hikingMap [][]int) int {
	sum := 0

	for y := 0; y < len(hikingMap); y++ {
		for x := 0; x < len(hikingMap[y]); x++ {
			if hikingMap[y][x] == 0 {
				sum += startWalking2(hikingMap, y, x)
			}
		}
	}

	return sum
}
func startWalking2(hikingMap [][]int, y int, x int) int {
	if hikingMap[y][x] == 9 {
		return 1
	}

	sum := 0
	next := hikingMap[y][x] + 1

	if y > 0 && hikingMap[y-1][x] == next {
		// walk up
		sum += startWalking2(hikingMap, y-1, x)
	}
	if y < len(hikingMap)-1 && hikingMap[y+1][x] == next {
		// walk down
		sum += startWalking2(hikingMap, y+1, x)
	}
	if x > 0 && hikingMap[y][x-1] == next {
		// walk left
		sum += startWalking2(hikingMap, y, x-1)
	}
	if x < len(hikingMap[y])-1 && hikingMap[y][x+1] == next {
		// walk right
		sum += startWalking2(hikingMap, y, x+1)
	}
	return sum
}
