package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

func main() {
	var data []string
	readFile(&data)

	a := part1(data)
	fmt.Printf("Part 1: %d\n", a)

	b := part2(data)
	fmt.Printf("Part 2: %d\n", b)
}

func calcSeatID(row int, col int) int {
	return (row * 8) + col
}

func getRow(line string, lower int, upper int) int {
	for i := 0; i < 7; i++ {
		char := string(line[i])
		var incr int = (upper - lower) / 2
		if char == "F" {
			upper = lower + incr
		} else if char == "B" {
			lower += incr + 1
		}
	}
	return lower
}

func getCol(line string, lower int, upper int) int {
	for i := 7; i < 10; i++ {
		char := string(line[i])
		var incr int = (upper - lower) / 2
		if char == "L" {
			upper = lower + incr
		} else if char == "R" {
			lower += incr + 1
		}
	}
	return lower
}

func getMax(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func getMin(x int, y int) int {
	if x < y {
		return x
	}
	return y
}

func findMissingSeat(min int, max int, m map[int]bool) int {
	for i := min + 1; i < max; i++ {
		prev := i - 1
		curr := i
		next := i + 1
		_, prevOk := m[prev]
		_, currOk := m[curr]
		_, nextOk := m[next]

		if !currOk && prevOk && nextOk {
			return curr
		}
	}
	return -1
}

func part2(data []string) int {
	minR := math.MaxInt32
	maxR := math.MinInt32
	for _, str := range data {
		row := getRow(str, 0, 127)
		maxR = getMax(maxR, row)
		minR = getMin(minR, row)
	}

	m := make(map[int]bool)
	min := math.MaxInt32
	max := math.MinInt32
	for _, str := range data {
		row := getRow(str, 0, 127)
		if row == minR || row == maxR {
			continue
		}

		seatId := calcSeatID(row, getCol(str, 0, 7))
		m[seatId] = true
		max = getMax(max, seatId)
		min = getMin(min, seatId)
	}
	return findMissingSeat(min, max, m)
}

func part1(data []string) int {
	max := 0
	for _, str := range data {
		seatId := calcSeatID(getRow(str, 0, 127), getCol(str, 0, 7))
		max = getMax(max, seatId)
	}
	return max
}

func readFile(data *[]string) {
	file, err := os.Open("./input/day5.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i := scanner.Text()
		*data = append(*data, i)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
