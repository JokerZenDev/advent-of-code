package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// input, err := os.ReadFile("test.txt")
	input, err := os.ReadFile("puzzle.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(input), "\n")

	startMap, startPos := createMap(lines)

	fmt.Println("first star:", firstStar(startMap, startPos))
}

var directions = [4]string{
	"up",
	"right",
	"down",
	"left",
}

func createMap(lines []string) (map[int]map[int]string, []int) {
	result := map[int]map[int]string{}
	startPos := []int{}
	for i, line := range lines {
		result[i] = map[int]string{}
		for j, char := range line {
			result[i][j] = string(char)
			if char == '^' {
				startPos = []int{i, j}
			}
		}
	}
	return result, startPos
}

func firstStar(startMap map[int]map[int]string, startPos []int) int {
	curPos := startPos

	dir := 0
	for curPos[0] < len(startMap) && curPos[0] >= 0 && curPos[1] < len(startMap[0]) && curPos[1] >= 0 {
		curPos = move(startMap, curPos, &dir)
		if dir > 3 {
			dir = 0
		}
	}

	sum := 0
	for _, row := range startMap {
		for _, col := range row {
			if col == "X" {
				sum++
			}
		}
	}

	return sum
}

func move(startMap map[int]map[int]string, curPos []int, direction *int) []int {
	dir := directions[*direction]

	oldPos := make([]int, len(curPos))
	copy(oldPos, curPos)

	switch dir {
	case "up":
		curPos[0]--
	case "down":
		curPos[0]++
	case "left":
		curPos[1]--
	case "right":
		curPos[1]++
	}

	if startMap[curPos[0]][curPos[1]] == "#" {
		*direction++
		return oldPos
	}

	startMap[oldPos[0]][oldPos[1]] = "X"
	return curPos
}
