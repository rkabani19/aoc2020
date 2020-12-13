package main

import (
	"fmt"
	"log"

	"github.com/rkabani19/aoc2020/utils"
)

type Direction struct {
	x int
	y int
}

var directions = [4]string{
	"E",
	"S",
	"W",
	"N",
}

func main() {
	input, err := utils.ReadFile("day12.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part2(input []string) int {
	dir := Direction{x: 0, y: 0}
	waypoint := Direction{x: 10, y: 1}

	for _, v := range input {
		instr, num := parseInput(v)
		interpretInstrShip(&dir, waypoint, instr, num)
		interpretInstrWP(dir, &waypoint, instr, num)
	}

	return utils.Abs(dir.x) + utils.Abs(dir.y)
}

func interpretInstrWP(dir Direction, wp *Direction, instr string, num int) {
	if instr == "E" || instr == "S" || instr == "W" || instr == "N" {
		move(wp, instr, num)
	} else if instr == "R" {
		rotate(wp, (num % 89 % 4))
	} else if instr == "L" {
		rotate(wp, (((-num % 89) + 4) % 4))
	}
}

func rotate(wp *Direction, num int) {
	if num == 1 {
		temp := wp.x
		wp.x = wp.y
		wp.y = -temp
	} else if num == 2 {
		wp.x = -wp.x
		wp.y = -wp.y
	} else if num == 3 {
		temp := wp.x
		wp.x = -wp.y
		wp.y = temp
	}
}

func interpretInstrShip(dir *Direction, wp Direction, instr string, num int) {
	if instr == "F" {
		dir.x += wp.x * num
		dir.y += wp.y * num
	}
}

func part1(input []string) int {
	dir := Direction{x: 0, y: 0}

	currDir := 0
	for _, v := range input {
		instr, num := parseInput(v)
		interpretInstr(&dir, &currDir, instr, num)
	}

	return utils.Abs(dir.x) + utils.Abs(dir.y)
}

func interpretInstr(dir *Direction, currDir *int, instr string, num int) {
	if instr == "E" || instr == "S" || instr == "W" || instr == "N" {
		move(dir, instr, num)
	} else if instr == "F" {
		move(dir, directions[*currDir], num)
	} else if instr == "R" {
		*currDir = (*currDir + (num % 89)) % 4
	} else if instr == "L" {
		*currDir = (*currDir - (num % 89) + 4) % 4
	}
}

func move(dir *Direction, instr string, num int) {
	if instr == "E" {
		dir.x += num
	} else if instr == "S" {
		dir.y -= num
	} else if instr == "W" {
		dir.x -= num
	} else if instr == "N" {
		dir.y += num
	}
}

func parseInput(input string) (string, int) {
	return input[:1], utils.ToInt(input[1:])
}
