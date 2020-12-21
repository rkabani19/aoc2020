package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/rkabani19/aoc2020/utils"
)

func main() {
	input, err := utils.ReadFile("day14.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part2(input []string) int64 {
	mem := make(map[int64]int64)
	currMask := ""

	for _, line := range input {
		currMask = updateMask(currMask, line)
		addr, val := parseWrite2(line)
		if addr == "" && val == 0 {
			continue
		}

		for i := 0; i < len(currMask); i++ {
			if currMask[i] != '0' {
				addr = utils.ReplaceAtIndex(addr, rune(currMask[i]), i)
			}
		}

		addresses := getPermutations(currMask, addr)
		for _, addrString := range addresses {
			addrInt, _ := strconv.ParseInt(addrString, 2, 64)
			mem[addrInt] = val
		}
	}

	return addMem(mem)
}

func getPermutations(currMask string, addr string) []string {
	res := make([]string, 0)
	masks := getMasks(strings.Count(currMask, "X"))

	for _, mask := range masks {
		cur := addr
		for i, j := 0, 0; i < len(cur); i++ {
			if cur[i] == 'X' {
				cur = utils.ReplaceAtIndex(cur, rune(mask[j]), i)
				j++
			}
		}
		res = append(res, cur)
	}
	return res
}

func getMasks(length int) []string {
	res := make([]string, 0)
	binStr := ""
	for i := 0; i < int(math.Pow(2, float64(length))); i++ {
		binStr = strconv.FormatInt(int64(i), 2)
		for len(binStr) < length {
			binStr = "0" + binStr
		}
		res = append(res, binStr)
	}
	return res
}

func parseWrite2(line string) (string, int64) {
	if line[0:4] == "mask" {
		return "", 0
	}

	split := strings.Split(line, " = ")

	addrStr := strings.Replace(strings.Split(split[0], "mem[")[1], "]", "", -1)
	addrStr = strconv.FormatInt(int64(utils.ToInt(addrStr)), 2)
	for len(addrStr) < 36 {
		addrStr = "0" + addrStr
	}

	val, _ := strconv.ParseInt(split[1], 10, 64)

	return addrStr, val
}

func part1(input []string) int64 {
	mem := make(map[int64]int64)
	currMask := ""

	for _, line := range input {
		currMask = updateMask(currMask, line)
		addr, val := parseWrite(line)
		if addr == 0 && val == "" {
			continue
		}

		for i := 0; i < len(currMask); i++ {
			if currMask[i] != 'X' {
				val = utils.ReplaceAtIndex(val, rune(currMask[i]), i)
			}
		}

		nVal, _ := strconv.ParseInt(val, 2, 64)
		mem[addr] = nVal
	}

	return addMem(mem)
}

func addMem(mem map[int64]int64) int64 {
	sum := int64(0)
	for _, v := range mem {
		sum += v
	}
	return sum
}

func parseWrite(line string) (int64, string) {
	if line[0:4] == "mask" {
		return 0, ""
	}

	split := strings.Split(line, " = ")

	addrStr := strings.Replace(strings.Split(split[0], "mem[")[1], "]", "", -1)
	addr, _ := strconv.ParseInt(addrStr, 10, 64)

	val := strconv.FormatInt(int64(utils.ToInt(split[1])), 2)
	for len(val) < 36 {
		val = "0" + val
	}

	return addr, val
}

func updateMask(curr string, line string) string {
	if line[0:4] == "mask" {
		return strings.Replace(line, "mask = ", "", -1)
	}
	return curr
}
