package main

import (
	"fmt"
	"strings"

	"github.com/rkabani19/aoc2020/utils"
)

type Tile struct {
	index        int
	edges        Edges
	flippedEdges Edges
}

type Edges struct {
	top    string
	bottom string
	left   string
	right  string
}

func main() {
	input, _ := utils.ReadFile("day20.txt")

	fmt.Printf("Part 1: %d\n", part1(input))
}

//func part2(input []string) int {
// NEED TO RETRY THIS
//}

func part1(input []string) int {
	tempTiles := parseInput(input)
	tiles := getTiles(tempTiles)
	corners := getCorners(tiles)

	res := 1
	for _, c := range corners {
		res *= c.index
	}
	return res
}

func getCorners(tiles []Tile) []Tile {
	corners := make([]Tile, 0)
	for i, t := range tiles {
		e := t.edges
		count := 0
		for j, t2 := range tiles {
			if i == j {
				continue
			}
			count += countMatchingEdges(e, t2)
		}
		if count == 2 {
			corners = append(corners, t)
		}
	}
	return corners
}

func countMatchingEdges(e Edges, t Tile) int {
	count := 0
	if edgesMatch(e.top, t) {
		count++
	}
	if edgesMatch(e.bottom, t) {
		count++
	}
	if edgesMatch(e.left, t) {
		count++
	}
	if edgesMatch(e.right, t) {
		count++
	}
	return count
}

func edgesMatch(e string, t Tile) bool {
	if e == t.edges.top || e == t.flippedEdges.top {
		return true
	}
	if e == t.edges.bottom || e == t.flippedEdges.bottom {
		return true
	}
	if e == t.edges.left || e == t.flippedEdges.left {
		return true
	}
	if e == t.edges.right || e == t.flippedEdges.right {
		return true
	}
	return false
}

func getTiles(m map[int][]string) []Tile {
	res := make([]Tile, 0)
	for k, v := range m {
		edges := getEdges(v)
		flippedEdges := flipEdges(edges)
		res = append(res, Tile{
			index:        k,
			edges:        edges,
			flippedEdges: flippedEdges,
		})
	}
	return res
}

func getEdges(t []string) Edges {
	e := Edges{}
	e.top = t[0]
	e.bottom = t[len(t)-1]

	left := make([]byte, 0)
	right := make([]byte, 0)
	for _, s := range t {
		left = append(left, s[0])
		right = append(right, s[len(s)-1])
	}
	e.left = string(left)
	e.right = string(right)

	return e
}

func flipEdges(e Edges) Edges {
	return Edges{
		top:    utils.ReverseString(e.top),
		bottom: utils.ReverseString(e.bottom),
		left:   utils.ReverseString(e.left),
		right:  utils.ReverseString(e.right),
	}
}

func parseInput(input []string) map[int][]string {
	res := make(map[int][]string)
	currTileNum := 0
	for _, line := range input {
		if strings.HasPrefix(line, "Tile") {
			sp := strings.Split(line, " ")
			currTileNum = utils.ToInt(string(sp[1][:len(sp[1])-1]))
			res[currTileNum] = make([]string, 0)
		} else if line == "" {
			continue
		} else {
			res[currTileNum] = append(res[currTileNum], line)
		}
	}
	return res
}
