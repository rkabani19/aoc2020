package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/rkabani19/aoc2020/utils"
)

func main() {
	input, err := utils.ReadFile("day13.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part2(input []string) int {
	_, ids := parseInput(input)
	sum, departure := utils.ToInt(ids[0]), utils.ToInt(ids[0]) // hard coded cause index 0 is not "x"

	// Chinese remainder theorum
	minApart := 0
	for i := 1; i < len(ids); i++ {
		minApart++
		if ids[i] == "x" {
			continue
		}
		id := utils.ToInt(ids[i])
		for (sum+minApart)%id != 0 {
			sum += departure
		}
		departure = (departure * id) / gcd(departure, id)
	}
	return sum
}

func gcd(x int, y int) int {
	div := 0
	for i := 1; i <= x && i <= y; i++ {
		if x%i == 0 && y%i == 0 {
			div = i
		}
	}
	return div
}

func part1(input []string) int {
	startTime, ids := parseInput(input)
	currTime := startTime
	for {
		for _, v := range ids {
			if v == "x" {
				continue
			}
			id := utils.ToInt(v)
			if currTime%id == 0 {
				return id * (currTime - startTime)
			}
		}
		currTime++
	}
}

func parseInput(input []string) (startTime int, ids []string) {
	startTime = utils.ToInt(input[0])
	ids = strings.Split(input[1], ",")
	return startTime, ids
}
