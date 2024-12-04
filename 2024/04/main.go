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

	fmt.Println("first star:", firstStar(string(input)))
	fmt.Println("second star:", secondStar(string(input)))
}

func firstStar(input string) int {
	sum := 0

	lines := strings.Split(input, "\n")
	height := len(lines)
	width := len(lines[0])
	diagonal := height + width
	allLines := make([]string, height+width+diagonal+diagonal)
	fmt.Println(height + width + diagonal)

	for pos, line := range lines {
		if len(line) < 4 {
			continue
		}
		allLines[pos] = line
		for i := 0; i < width; i++ {
			char := string(line[i])
			allLines[height+i] += char
			allLines[height+width+int(diagonal/2)+(i-pos)] += char
			allLines[height+width+diagonal+(i+pos)] += char
		}
	}
	checkLines(allLines, &sum)

	return sum
}

func checkLines(lines []string, sum *int) {
	for _, line := range lines {
		if len(line) < 4 {
			continue
		}
		for i := 0; i <= len(line)-4; i++ {
			if line[i:i+4] == "XMAS" || line[i:i+4] == "SAMX" {
				*sum++
			}
		}
	}
}

func secondStar(input string) int {
	sum := 0

	lines := strings.Split(input, "\n")

	for c := 1; c < len(lines)-1; c++ {
		line := lines[c]
		for i := 1; i < len(line)-1; i++ {
			char := string(line[i])
			if char == "A" {
				topLeft := string(lines[c-1][i-1])
				topRight := string(lines[c-1][i+1])
				bottomLeft := string(lines[c+1][i-1])
				bottomRight := string(lines[c+1][i+1])
				leftToRight := string(topLeft) + string(line[i]) + string(bottomRight)
				rightToLeft := string(topRight) + string(line[i]) + string(bottomLeft)
				if (leftToRight == "MAS" || leftToRight == "SAM") && (rightToLeft == "MAS" || rightToLeft == "SAM") {
					sum++
				}
			}
		}
	}

	return sum
}
