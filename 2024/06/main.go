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

	// fmt.Println("first star:", firstStar(startMap, startPos))
	fmt.Println("second star:", secondStar(startMap, startPos))
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

func secondStar(startMap map[int]map[int]string, startPos []int) int {
	curPos := make([]int, len(startPos))
	copy(curPos, startPos)

	dir := 0
	steps := [][]int{}

	curMap := copyMap(startMap)
	for curPos[0] < len(curMap) && curPos[0] >= 0 && curPos[1] < len(curMap[0]) && curPos[1] >= 0 {
		steps = append(steps, []int{curPos[0], curPos[1]})
		curPos = move(curMap, curPos, &dir)
		if dir > 3 {
			dir = 0
		}
	}

	// Remove duplicate steps
	uniqueSteps := [][]int{}
	seen := make(map[string]bool)

	for _, step := range steps {
		key := fmt.Sprintf("%d,%d", step[0], step[1])
		if !seen[key] {
			seen[key] = true
			uniqueSteps = append(uniqueSteps, step)
		}
	}
	steps = uniqueSteps

	loopsFound := 0
	for i := 1; i < len(steps); i++ {
		step := steps[i]
		dir = 0
		findLoop := false
		curMap = copyMap(startMap)
		curMap[step[0]][step[1]] = "0"
		copy(curPos, startPos)
		for curPos[0] < len(curMap) && curPos[0] >= 0 && curPos[1] < len(curMap[0]) && curPos[1] >= 0 {
			curPos, findLoop = move2(curMap, curPos, &dir)
			if findLoop {
				break
			}
			if dir > 3 {
				dir = 0
			}
		}
		if findLoop {
			loopsFound++
		}
	}

	return loopsFound
}

func move2(startMap map[int]map[int]string, curPos []int, direction *int) ([]int, bool) {
	dir := directions[*direction]
	findLoop := false

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

	if startMap[curPos[0]][curPos[1]] == "#" || startMap[curPos[0]][curPos[1]] == "0" {
		// startMap[oldPos[0]][oldPos[1]] = "+"
		*direction++
		return oldPos, findLoop
	}

	if startMap[oldPos[0]][oldPos[1]] == "." {
		startMap[oldPos[0]][oldPos[1]] = ""
	}

	switch dir {
	case "up":
		if strings.Contains(startMap[oldPos[0]][oldPos[1]], "U") {
			findLoop = true
		} else {
			startMap[oldPos[0]][oldPos[1]] += "U"
		}
	case "down":
		if strings.Contains(startMap[oldPos[0]][oldPos[1]], "D") {
			findLoop = true
		} else {
			startMap[oldPos[0]][oldPos[1]] += "D"
		}
	case "left":
		if strings.Contains(startMap[oldPos[0]][oldPos[1]], "L") {
			findLoop = true
		} else {
			startMap[oldPos[0]][oldPos[1]] += "L"
		}
	case "right":
		if strings.Contains(startMap[oldPos[0]][oldPos[1]], "R") {
			findLoop = true
		} else {
			startMap[oldPos[0]][oldPos[1]] += "R"
		}
	}

	return curPos, findLoop
}

func printMap(startMap map[int]map[int]string) {
	for i := 0; i < len(startMap); i++ {
		str := ""
		for j := 0; j < len(startMap[i]); j++ {
			str += startMap[i][j]
		}
		fmt.Println(str)
	}
	fmt.Println()
}

func copyMap(startMap map[int]map[int]string) map[int]map[int]string {
	curMap := make(map[int]map[int]string, len(startMap))
	for k, v := range startMap {
		curMap[k] = make(map[int]string, len(v))
		for key, val := range v {
			curMap[k][key] = val
		}
	}
	return curMap
}
