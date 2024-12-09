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
	antennasMap := [][]string{}
	for _, line := range lines {
		antennasMap = append(antennasMap, strings.Split(line, ""))
	}

	fmt.Println("first star:", firstStar(antennasMap))
}

func firstStar(antennasMap [][]string) int {
	antinodes := [][2]int{}

	for i := 0; i < len(antennasMap); i++ {
		for j := 0; j < len(antennasMap[i]); j++ {
			if antennasMap[i][j] != "." {
				posFound := [][2]int{{i, j}}
				sameFrequencyY, sameFrequencyX := findSameFrequency(antennasMap, antennasMap[i][j], posFound)
				for sameFrequencyY != -1 && sameFrequencyX != -1 {
					posFound = append(posFound, [2]int{sameFrequencyY, sameFrequencyX})
					distanceY := sameFrequencyY - i
					distanceX := sameFrequencyX - j

					antinodePosY := sameFrequencyY + distanceY
					antinodePosX := sameFrequencyX + distanceX

					sameFrequencyY, sameFrequencyX = findSameFrequency(antennasMap, antennasMap[i][j], posFound)
					if antinodePosY < 0 || antinodePosY >= len(antennasMap) || antinodePosX < 0 || antinodePosX >= len(antennasMap[antinodePosY]) {
						continue
					}

					notFound := true
					for _, antinode := range antinodes {
						if antinode[0] == antinodePosY && antinode[1] == antinodePosX {
							notFound = false
							break
						}
					}

					if notFound {
						antinodes = append(antinodes, [2]int{antinodePosY, antinodePosX})
					}
				}
			}
		}
	}

	return len(antinodes)
}

func findSameFrequency(antennasMap [][]string, frequency string, posFound [][2]int) (int, int) {
	for i := 0; i < len(antennasMap); i++ {
		for j := 0; j < len(antennasMap[i]); j++ {
			if antennasMap[i][j] == frequency {
				notFound := true
				for _, pos := range posFound {
					if pos[0] == i && pos[1] == j {
						notFound = false
						break
					}
				}
				if notFound {
					return i, j
				}
			}
		}
	}
	return -1, -1
}
