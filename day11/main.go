package main

import (
	"fmt"
	"log"
	"reflect"

	"github.com/rkabani19/aoc2020/utils"
)

var Locations = [8][2]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{1, 1},
	{1, -1},
	{-1, 0},
	{-1, 1},
	{-1, -1},
}

func main() {
	input, err := utils.ReadFile("day11.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part2(input []string) int {
	curr := make([]string, len(input))
	copy(curr, input)
	prev := make([]string, len(input))
	copy(prev, input)

	for {
		fillSeats(curr)
		if reflect.DeepEqual(curr, prev) {
			break
		}
		copy(prev, curr)
	}

	return numOccupied(curr)
}

func part1(input []string) int {
	curr := make([]string, len(input))
	copy(curr, input)
	prev := make([]string, len(input))
	copy(prev, input)

	for {
		fillSeats1(curr)
		if reflect.DeepEqual(curr, prev) {
			break
		}
		copy(prev, curr)
	}

	return numOccupied(curr)
}

func fillSeats(input []string) {
	currInput := make([]string, len(input))
	copy(currInput, input)
	for r := 0; r < len(input); r++ {
		line := input[r]
		for c := 0; c < len(line); c++ {
			numAdj := getNumAdj(input, r, c)
			curr := input[r][c]
			if curr == 'L' && numAdj == 0 {
				currInput[r] = replaceAtIndex(currInput[r], '#', c)
			} else if curr == '#' && numAdj >= 5 {
				currInput[r] = replaceAtIndex(currInput[r], 'L', c)
			}
		}
	}
	copy(input, currInput)
}

func getNumAdj(input []string, row int, col int) int {
	numAdj := 0
	for _, a := range Locations {
		dx, dy := a[0], a[1]
		nCol, nRow := col+dx, row+dy
		for nRow >= 0 && nRow < len(input) && nCol >= 0 && nCol < len(input[nRow]) {
			if input[nRow][nCol] == '#' {
				numAdj++
				break
			}
			if input[nRow][nCol] == 'L' {
				break
			}
			nCol += dx
			nRow += dy
		}
	}
	return numAdj
}

// For Part 1
func fillSeats1(input []string) {
	currInput := make([]string, len(input))
	copy(currInput, input)
	for r := 0; r < len(input); r++ {
		line := input[r]
		for c := 0; c < len(line); c++ {
			numAdj := getNumAdj1(input, r, c)
			curr := input[r][c]
			if curr == 'L' && numAdj == 0 {
				currInput[r] = replaceAtIndex(currInput[r], '#', c)
			} else if curr == '#' && numAdj >= 4 {
				currInput[r] = replaceAtIndex(currInput[r], 'L', c)
			}
		}
	}
	copy(input, currInput)
}

func getNumAdj1(input []string, row int, col int) int {
	numAdj := 0
	for _, a := range Locations {
		dx, dy := a[0], a[1]
		nCol, nRow := col+dx, row+dy
		if nRow < 0 || nRow >= len(input) || nCol < 0 || nCol >= len(input[row]) {
			continue
		}
		if input[nRow][nCol] == '#' {
			numAdj++
		}
	}
	return numAdj
}

func numOccupied(input []string) int {
	num := 0
	for r := 0; r < len(input); r++ {
		for c := 0; c < len(input[r]); c++ {
			if input[r][c] == '#' {
				num++
			}
		}
	}
	return num
}

func replaceAtIndex(str string, replacement rune, index int) string {
	out := []rune(str)
	out[index] = replacement
	return string(out)
}
