package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/rkabani19/aoc2020/utils"
)

func main() {
	data, err := utils.ReadFile("day8.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %d\n", part1(data))
	fmt.Printf("Part 2: %d\n", part2(data))
}

func part2(data []string) int {
	switches := make(map[int]bool)

	for {
		visited := make(map[int]bool)
		instrSwitched := false
		valid := true
		accum := 0

		for i := 0; i < len(data); i++ {
			instr, num := parseInstr(data[i])

			if _, ok := switches[i]; !ok && !instrSwitched && instr != "acc" {
				if instr == "jmp" {
					instr = "nop"
				} else if instr == "nop" {
					instr = "jmp"
				}
				instrSwitched = true
				switches[i] = true
			}

			if _, ok := visited[i]; ok {
				valid = false
				break
			}

			visited[i] = true
			addI, addAcc := handleInstr(instr, num)
			i += addI
			accum += addAcc
		}

		if valid {
			return accum
		}
	}
}

func part1(data []string) int {
	visited := make(map[int]bool)
	accum := 0
	for i := 0; i < len(data); i++ {
		instr, num := parseInstr(data[i])
		if _, ok := visited[i]; ok {
			break
		}

		visited[i] = true

		addI, addAcc := handleInstr(instr, num)
		i += addI
		accum += addAcc
	}

	return accum
}

func parseInstr(instr string) (string, int) {
	// should've used a struct, oh well D:
	split := strings.Split(instr, " ")
	rInstr := split[0]
	rNum, _ := strconv.Atoi(split[1])
	return rInstr, rNum
}

func handleInstr(instr string, num int) (int, int) {
	if instr == "acc" {
		return 0, num
	} else if instr == "jmp" {
		return num - 1, 0
	}
	return 0, 0
}
