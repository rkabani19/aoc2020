package main

import (
	"fmt"

	"github.com/rkabani19/aoc2020/utils"
)

type queue []int

func main() {
	input, _ := utils.ReadFile("day22.txt")

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part2(input []string) int {
	p1, p2 := parseInput(input)
	p1, p2 = simulateRecur(p1, p2)
	if len(p1) == 0 {
		return calculateScore(p2)
	}
	return calculateScore(p1)
}

func simulateRecur(p1 queue, p2 queue) (queue, queue) {
	pastStates := make(map[string]bool)
	for len(p1) != 0 && len(p2) != 0 {
		k := fmt.Sprintf("%v%v", p1, p2)
		if pastStates[k] {
			return append(p1, p2...), nil
		}
		pastStates[k] = true

		c1, c2 := p1[0], p2[0]

		//	fmt.Printf("Player 1's deck: %v\n", p1)
		//	fmt.Printf("Player 2's deck: %v\n", p2)

		p1, p2 = p1[1:], p2[1:]

		//	fmt.Printf("Player 1 plays: %v\n", c1)
		//	fmt.Printf("Player 2 plays: %v\n", c2)

		if c1 <= len(p1) && c2 <= len(p2) {
			//		fmt.Printf("Playing a subgame now...\n\n")
			t1, t2 := simulateRecur(cp(p1[:c1]), cp(p2[:c2]))

			if len(t1) == 0 {
				p2 = append(p2, c2, c1)
			} else if len(t2) == 0 {
				p1 = append(p1, c1, c2)
			} else {
				//			fmt.Println("Error in recur game.")
			}
		} else {
			if c1 > c2 {
				p1 = append(p1, c1, c2)
				//			fmt.Printf("Player 1 wins.\n\n")
			} else {
				p2 = append(p2, c2, c1)
				//			fmt.Printf("Player 2 wins.\n\n")
			}
		}
	}

	return p1, p2
}

func cp(v []int) []int {
	return append([]int(nil), v...)
}

func part1(input []string) int {
	p1, p2 := parseInput(input)
	p1, p2 = simulate(p1, p2)
	if len(p1) == 0 {
		return calculateScore(p2)
	}
	return calculateScore(p1)
}

func calculateScore(q queue) int {
	score := 0
	sz := len(q)
	for _, n := range q {
		score += sz * n
		sz--
	}
	return score
}

func simulate(p1 queue, p2 queue) (queue, queue) {
	for len(p1) != 0 && len(p2) != 0 {
		c1, c2 := p1[0], p2[0]
		p1 = p1[1:]
		p2 = p2[1:]
		if c1 > c2 {
			p1 = append(p1, c1, c2)
		} else {
			p2 = append(p2, c2, c1)
		}
	}
	return p1, p2
}

func parseInput(input []string) (p1 queue, p2 queue) {
	i := 1
	p1, p2 = make(queue, 0), make(queue, 0)
	for input[i] != "" {
		p1 = append(p1, utils.ToInt(input[i]))
		i++
	}
	i += 2
	for i < len(input) {
		p2 = append(p2, utils.ToInt(input[i]))
		i++
	}
	return p1, p2
}
