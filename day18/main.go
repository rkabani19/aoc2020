package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rkabani19/aoc2020/utils"
)

func main() {
	input, _ := utils.ReadFile("day18.txt")

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part2(input []string) int {
	count := 0
	for _, line := range input {
		count += calculate2([]byte(strings.Replace(line, " ", "", -1)))
	}
	return count
}

func calculate2(expr []byte) int {
	var s utils.Stack
	res1, res2 := 0, 1
	for _, b := range expr {
		if b == '(' {
			s.Push(strconv.Itoa(res1))
			s.Push(strconv.Itoa(res2))

			res1, res2 = 0, 1
		} else if b == ')' {
			res := res1 * res2

			strres2, _ := s.Pop()
			res2 = utils.ToInt(strres2)
			strres1, _ := s.Pop()
			res1 = utils.ToInt(strres1)

			res1 += res
		} else if b == '*' {
			res2 = res1 * res2
			res1 = 0
		} else if b == '+' {
		} else {
			res1 += utils.ToInt(string(b))
		}
	}
	return res1 * res2
}

func part1(input []string) int {
	count := 0
	for _, line := range input {
		count += calculate([]byte(strings.Replace(line, " ", "", -1)))
	}
	return count
}

func calculate(expr []byte) int {
	var s utils.Stack
	op := "+"
	res := 0
	for _, b := range expr {
		if b == '+' || b == '*' {
			op = string(b)
		} else if b == '(' {
			s.Push(strconv.Itoa(res))
			s.Push(op)
			res = 0
			op = "+"
		} else if b == ')' {
			op, _ = s.Pop()
			prevRes, _ := s.Pop()
			res = operate(utils.ToInt(prevRes), op, res)
		} else {
			res = operate(res, op, utils.ToInt(string(b)))
		}
	}
	return res
}

func operate(x int, op string, y int) int {
	if op == "*" {
		return x * y
	}
	return x + y
}
