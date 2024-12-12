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
	stones := strings.Split(string(input), " ")

	fmt.Println("first star:", firstStar(stones))
}

func firstStar(stones []string) int {
	blinks := 25

	for i := 0; i < blinks; i++ {
		stones = loopStones(stones)
	}

	return len(stones)
}

func loopStones(stones []string) []string {
	newStones := []string{}
	for i := 0; i < len(stones); i++ {
		newStones = append(newStones, blinkChange(stones[i])...)
	}
	return newStones
}

var cache = map[string][]string{}

func blinkChange(stone string) []string {
	cached, isCached := cache[stone]
	if isCached {
		return cached
	}
	if stone == "0" {
		newStones := []string{"1"}
		cache[stone] = newStones
		return newStones
	}
	if len(stone)%2 == 0 {
		leftHalf, _ := strconv.Atoi(stone[:len(stone)/2])
		rightHalf, _ := strconv.Atoi(stone[len(stone)/2:])
		newStones := []string{strconv.Itoa(leftHalf)}
		newStones = append(newStones, strconv.Itoa(rightHalf))
		cache[stone] = newStones
		return newStones
	}
	num, _ := strconv.Atoi(stone)
	newStones := []string{strconv.Itoa(num * 2024)}
	cache[stone] = newStones
	return newStones
}

