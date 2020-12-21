package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rkabani19/aoc2020/utils"
)

func main() {
	input, _ := utils.ReadFile("day15.txt")

	fmt.Printf("Part 1: %d\n", part1and2(input, 2020))
	fmt.Printf("Part 2: %d\n", part1and2(input, 30000000))
}

func part1and2(input []string, itr int) int {
	m := make(map[int][]int)
	nInput := strings.Split(input[0], ",")
	var prev int

	i := 0
	for i < len(nInput) {
		num, _ := strconv.ParseInt(string(nInput[i]), 10, 64)
		m[int(num)] = []int{i + 1, -1}
		prev = int(num)
		i++
	}

	for i < itr {
		spoken := 0
		if v, ok := m[prev]; ok {
			if v[1] != -1 {
				spoken = v[0] - v[1]
			}
		}

		if v, ok := m[spoken]; ok {
			v[1] = v[0]
			v[0] = i + 1
		} else {
			m[spoken] = []int{i + 1, -1}
		}

		prev = spoken
		i++
	}

	return prev
}
