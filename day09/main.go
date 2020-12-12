package main

import (
	"fmt"
	"log"
	"math"
	"strconv"

	"github.com/rkabani19/aoc2020/utils"
)

const PreambleLength = 25

func main() {
	data, err := utils.ReadFile("day9.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %d\n", part1(data))
	fmt.Printf("Part 2: %d\n", part2(data))
}

func part2(input []string) int {
	target := part1(input)
	nInput := parseInputToIntArr(input)

	start, end := findContiguousSum(nInput, target)

	return minInArray(nInput, start, end) + maxInArray(nInput, start, end)
}

func findContiguousSum(input []int, target int) (int, int) {
	sumArr := make([]int, len(input))
	sumArr[0] = input[0]
	for i := 1; i < len(sumArr); i++ {
		sumArr[i] = input[i] + sumArr[i-1]
	}

	for i := 0; i < len(input); i++ {
		for j := i; j < len(input); j++ {
			diff := sumArr[j] - sumArr[i]
			if diff == target {
				return i, j
			}
		}
	}

	return -1, -1
}

func minInArray(input []int, start int, end int) int {
	min := math.MaxInt32
	for i := start; i < end; i++ {
		if min > input[i] {
			min = input[i]
		}
	}
	return min
}

func maxInArray(input []int, start int, end int) int {
	max := math.MinInt32
	for i := start; i < end; i++ {
		if max < input[i] {
			max = input[i]
		}
	}
	return max
}

func part1(input []string) int {
	nInput := parseInputToIntArr(input)

	for i := PreambleLength; i < len(nInput); i++ {
		num := nInput[i]
		if !hasSum(nInput, i, num) {
			return num
		}
	}

	return -1
}

func hasSum(arr []int, end int, target int) bool {
	m := make(map[int]bool)
	for i := end - PreambleLength; i < end; i++ {
		num := arr[i]
		diff := target - num
		if _, ok := m[num]; ok {
			return true
		}
		m[diff] = true
	}
	return false
}

func parseInputToIntArr(input []string) []int {
	res := make([]int, len(input))
	for i := 0; i < len(input); i++ {
		num, _ := strconv.Atoi(input[i])
		res[i] = num
	}
	return res
}
