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
	numbers := strings.Split(string(input), "")
	numbersInt := []int{}
	for _, number := range numbers {
		num, err := strconv.Atoi(number)
		if err != nil {
			panic(err)
		}
		numbersInt = append(numbersInt, num)
	}

	fmt.Println("first star:", firstStar(numbersInt))
}

func firstStar(numbers []int) int {
	disk := createDisk(numbers)
	disk = moveIndividual(disk)
	sum := checksum(disk)

	return sum
}

func createDisk(numbers []int) []int {
	disk := []int{}
	number := 0

	for i := 0; i < len(numbers); i++ {
		for c := 0; c < numbers[i]; c++ {
			if i%2 == 0 {
				disk = append(disk, number)
			} else {
				disk = append(disk, -1)
			}
		}
		if i%2 == 0 {
			number++
		}
	}

	return disk
}

func moveIndividual(disk []int) []int {
	for i := len(disk) - 1; i >= 0; i-- {
		if disk[i] != -1 {
			exchange := false
			for j := 0; j < i; j++ {
				if disk[j] == -1 {
					disk[j] = disk[i]
					disk[i] = -1
					exchange = true
					break
				}
			}
			if !exchange {
				break
			}
		}
	}

	return disk
}

func checksum(disk []int) int {
	sum := 0
	for i, num := range disk {
		if num != -1 {
			sum += num * i
		}
	}
	return sum
}
