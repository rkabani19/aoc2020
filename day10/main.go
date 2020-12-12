package main

import (
	"fmt"
	"log"
	"math"

	"github.com/rkabani19/aoc2020/utils"
)

const MaxDiff = 3

func main() {
	input, err := utils.ReadFile("day10.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part2(input []string) int {
	m, max, min := getIntMap(input)
	if min > 3 {
		fmt.Println("Cannot find suitable adapter to start.")
		return -1
	}

	memo := make(map[int]int) // k: curr val, v: result
	return getDistinctArrangements(m, memo, max, 1) +
		getDistinctArrangements(m, memo, max, 2) +
		getDistinctArrangements(m, memo, max, 3)
}

func getDistinctArrangements(m map[int]bool, memo map[int]int, max int, curr int) int {
	if _, ok := m[curr]; !ok {
		return 0
	}
	if max == curr {
		return 1
	}
	if v, ok := memo[curr]; ok {
		return v
	}

	res := getDistinctArrangements(m, memo, max, curr+1) +
		getDistinctArrangements(m, memo, max, curr+2) +
		getDistinctArrangements(m, memo, max, curr+3)

	memo[curr] = res
	return res
}

func part1(input []string) int {
	prev := 0
	count1, count3 := 0, 0
	m, max, min := getIntMap(input)
	for i := min; i <= max; i++ {
		_, ok := m[i]
		if !ok {
			continue
		}

		diff := i - prev
		if diff == 1 {
			count1++
		} else if diff == 2 {
		} else if diff == 3 {
			count3++
		} else {
			fmt.Println("Adapters cannot be linked.")
			break
		}

		prev = i
	}

	return count1 * (count3 + 1)
}

func getIntMap(input []string) (map[int]bool, int, int) {
	m := make(map[int]bool)
	max, min := math.MinInt32, math.MaxInt32
	for _, str := range input {
		num := utils.ToInt(str)
		m[num] = true
		max = utils.GetMax(max, num)
		min = utils.GetMin(min, num)
	}
	return m, max, min
}
