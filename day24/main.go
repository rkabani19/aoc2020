package main

import (
	"fmt"

	"github.com/rkabani19/aoc2020/utils"
)

type Direction int
type Hex [3]Direction
type Color int
type Grid map[Hex]Color

const (
	E Direction = iota
	SE
	SW
	W
	NW
	NE
)

const (
	White = iota
	Black
)

func (h Hex) walk(d Direction) Hex {
	if d == E {
		return Hex{h[0] + 1, h[1], h[2] - 1}
	} else if d == SE {
		return Hex{h[0], h[1] + 1, h[2] - 1}
	} else if d == SW {
		return Hex{h[0] - 1, h[1] + 1, h[2]}
	} else if d == W {
		return Hex{h[0] - 1, h[1], h[2] + 1}
	} else if d == NW {
		return Hex{h[0], h[1] - 1, h[2] + 1}
	} else if d == NE {
		return Hex{h[0] + 1, h[1] - 1, h[2]}
	}
	panic(fmt.Errorf("No valid direction."))
}

func main() {
	input, _ := utils.ReadFile("day24.txt")

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part2(input []string) int {
	tiles := parseInput(input)

	g := make(Grid)
	for _, t := range tiles {
		g.flipTile(t)
	}

	for i := 0; i < 100; i++ {
		g = g.dailyFlip()
		//fmt.Printf("Day %d: %d\n", i+1, g.countBlackTiles())
	}

	return g.countBlackTiles()
}

func (g Grid) dailyFlip() Grid {
	out := make(Grid)
	for h := range g {
		for i := Direction(0); i < 6; i++ {
			neighbor := h.walk(i)
			countBlack := 0
			for j := Direction(0); j < 6; j++ {
				if g[neighbor.walk(j)] == Black {
					countBlack++
				}
			}
			if g[neighbor] == Black && countBlack != 0 && countBlack <= 2 {
				out[neighbor] = Black
			} else if g[neighbor] == White && countBlack == 2 {
				out[neighbor] = Black
			}
		}
	}
	return out
}

func part1(input []string) int {
	tiles := parseInput(input)

	g := make(Grid)
	for _, t := range tiles {
		g.flipTile(t)
	}

	return g.countBlackTiles()
}

func (g Grid) countBlackTiles() int {
	count := 0
	for _, v := range g {
		if v == Black {
			count++
		}
	}
	return count
}

func (g Grid) flipTile(dirs []Direction) {
	h := Hex{}
	for _, d := range dirs {
		h = h.walk(d)
	}

	if g[h] == Black {
		g[h] = White
	} else {
		g[h] = Black
	}
}

func parseInput(input []string) [][]Direction {
	out := make([][]Direction, 0)
	for _, line := range input {
		d := make([]Direction, 0)
		for len(line) > 0 {
			l := line[0]
			if l == 'e' {
				d = append(d, E)
				line = line[1:]
			} else if l == 'w' {
				d = append(d, W)
				line = line[1:]
			} else if l == 'n' {
				if line[1] == 'e' {
					d = append(d, NE)
				} else {
					d = append(d, NW)
				}
				line = line[2:]
			} else if l == 's' {
				if line[1] == 'e' {
					d = append(d, SE)
				} else {
					d = append(d, SW)
				}
				line = line[2:]
			}
		}
		out = append(out, d)
	}
	return out
}
