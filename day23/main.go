package main

import (
	"fmt"
	"strconv"

	"github.com/rkabani19/aoc2020/utils"
)

type circle []int

type Game struct {
	cur  int
	cups circle
}

func main() {
	input, _ := utils.ReadFile("day23.txt")

	//fmt.Printf("Part 1: %s\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part2(input []string) int {
	cups := parseInput(input)
	for i := len(cups) + 1; i <= 1000000; i++ {
		cups = append(cups, i)
	}
	g := Game{cur: 0, cups: cups}

	for i := 0; i < 10000000; i++ {
		g = doMove(g)
	}

	i := 0
	for g.cups[i] != 1 {
		i++
	}

	return g.cups[i+1] * g.cups[i+2]
}

func part1(input []string) string {
	cups := parseInput(input)
	g := Game{cur: 0, cups: cups}

	for i := 0; i < 100; i++ {
		g = doMove(g)
	}

	return orderCups(g.cups)
}

func orderCups(cups circle) string {
	out := ""
	start := cups.find(1)
	for i := start + 1; i < len(cups); i++ {
		out = out + strconv.Itoa(cups[i])
	}
	for i := 0; cups[i] != 1; i++ {
		out = out + strconv.Itoa(cups[i])
	}
	return out
}

func doMove(g Game) Game {
	//fmt.Println("-- move --")

	cur := g.cups[g.cur]
	//fmt.Printf("Current: %d\n", cur)

	pickUp, remaining := g.cups.slice(g.cur+1, g.cur+4), g.cups.slice(g.cur+4, g.cur+1)
	//fmt.Printf("Pickup: %v\n", pickUp)

	dest := getDest(cur, pickUp)
	//fmt.Printf("Destination: %d\n", dest)

	d := remaining.find(dest)

	nc := make(circle, 0)
	nc = append(nc, dest)
	nc = append(nc, pickUp...)
	nc = append(nc, remaining.slice(d+1, d)...)

	//fmt.Println("")
	return Game{
		cur:  (nc.find(cur) + 1) % len(nc),
		cups: nc,
	}
}

func getDest(cur int, pickUp circle) int {
	des := cur - 1
	if des < 1 {
		des = 1000000 // 9 for part1
	}
	for pickUp.find(des) != -1 {
		des--
		if des < 1 {
			des = 1000000 // 9 for part1
		}
	}
	return des
}

func (c circle) shift(n int) circle {
	n = ((n % len(c)) + len(c)) % len(c)
	return append(append(circle(nil), c[:n]...), c[:n]...)
}

func (c circle) find(n int) int {
	for i, v := range c {
		if n == v {
			return i
		}
	}
	return -1
}

func (c circle) slice(i int, j int) circle {
	i, j = i%len(c), j%len(c)
	if j >= i {
		return c[i:j]
	}
	return append(append(circle(nil), c[i:]...), c[:j]...)
}

func parseInput(input []string) circle {
	out := make([]int, 0)
	for i := 0; i < len(input[0]); i++ {
		out = append(out, utils.ToInt(string(input[0][i])))
	}
	return out
}
