package main

import (
	"fmt"

	"github.com/rkabani19/aoc2020/utils"
)

const cycles = 6

type cell [3]int // x, y, z
type grid map[cell]bool

func main() {
	input, _ := utils.ReadFile("day17.txt")
	grid := parseGrid(input)

	fmt.Printf("Part 1: %d\n", part1(grid))
}

func part1(g grid) int {
	for i := 0; i < cycles; i++ {
		g = g.simulate()
	}
	return len(g)
}

func (g grid) simulate() grid {
	ng := make(grid)
	for c := range g {
		if count := g.countAliveNeighbors(c); count == 2 || count == 3 {
			ng[c] = true
		}
		g.walkNeighbors(c, func(c cell) {
			if g[c] {
				return
			}
			if count := g.countAliveNeighbors(c); count == 3 {
				ng[c] = true

			}
		})
	}
	return ng
}

func (g grid) countAliveNeighbors(c cell) int {
	count := 0
	g.walkNeighbors(c, func(c cell) {
		if g[c] {
			count++
		}
	})
	return count
}

func (g grid) walkNeighbors(c cell, f func(c cell)) {
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			for z := -1; z <= 1; z++ {
				if x == 0 && y == 0 && z == 0 {
					continue
				}
				curr := cell{c[0] + x, c[1] + y, c[2] + z}
				f(curr)
			}
		}
	}
}

func parseGrid(input []string) grid {
	grid := make(grid)
	for y, line := range input {
		for x, v := range []byte(line) {
			if v == '#' {
				grid[cell{x, y, 0}] = true
			}
		}
	}
	return grid
}
