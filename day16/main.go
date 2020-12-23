package main

import (
	"fmt"
	"strings"

	"github.com/rkabani19/aoc2020/utils"
)

type field struct {
	name string
	a    int
	b    int
	c    int
	d    int
}

func main() {
	input, _ := utils.ReadFile("day16.txt")
	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part2(input []string) int {
	ranges, myTicket, nearbyTickets := parseInput(input)

	validTickets := getValidTickets(nearbyTickets, ranges)

	possibleLocations := getPossibleLocations(ranges, validTickets)
	positions := getPositions(possibleLocations)

	product := 1
	for k, v := range positions {
		if strings.HasPrefix(k, "departure") {
			product *= myTicket[v]
		}
	}

	return product
}

func getPositions(possibleLocations map[string]map[int]bool) map[string]int {
	res := make(map[string]int)
	for len(possibleLocations) > 0 {
		for i, locations := range possibleLocations {
			if len(locations) == 1 {
				location := 0
				for location = range locations {
				}
				res[i] = location

				for k := range possibleLocations {
					delete(possibleLocations[k], location)
				}

				delete(possibleLocations, i)
				break
			}
		}
	}
	return res
}

func getPossibleLocations(ranges []field, tickets [][]int) map[string]map[int]bool {
	possibleLocations := make(map[string]map[int]bool)
	for _, r := range ranges {
		if _, ok := possibleLocations[r.name]; !ok {
			possibleLocations[r.name] = make(map[int]bool)
		}
	search:
		for i := 0; i < len(tickets[0]); i++ {
			for _, ticket := range tickets {
				if !isValid(r, ticket[i]) {
					continue search
				}
			}
			possibleLocations[r.name][i] = true
		}
	}

	return possibleLocations
}

func getValidTickets(nearbyTickets [][]int, ranges []field) [][]int {
	validTickets := make([][]int, 0)

search:
	for _, ticket := range nearbyTickets {
	check:
		for _, num := range ticket {
			for _, r := range ranges {
				if isValid(r, num) {
					continue check
				}
			}
			continue search
		}
		validTickets = append(validTickets, ticket)
	}

	return validTickets
}

func part1(input []string) int {
	ranges, _, nearbyTickets := parseInput(input)
	sum := 0

	for _, ticket := range nearbyTickets {
		for _, num := range ticket {
			valid := true
			for _, v := range ranges {
				if !isValid(v, num) {
					valid = false
				} else {
					valid = true
					break
				}
			}
			if !valid {
				sum += num
			}
		}
	}

	return sum
}

func isValid(f field, num int) bool {
	return (num >= f.a && num <= f.b) || (num >= f.c && num <= f.d)
}

func parseInput(input []string) (ranges []field, ticket []int, nearbyTickets [][]int) {
	i := 0
	ranges = getRanges(input, &i)
	i += 2
	ticket = getTicket(input, i)
	i += 3
	nearbyTickets = getNearbyTickets(input, i)
	return ranges, ticket, nearbyTickets
}

func getNearbyTickets(input []string, i int) [][]int {
	nearbyTickets := make([][]int, 0)
	for ; i < len(input); i++ {
		nearbyTickets = append(nearbyTickets, getTicket(input, i))
	}
	return nearbyTickets
}

func getTicket(input []string, i int) []int {
	strRet := strings.Split(input[i], ",")
	intRet := make([]int, 0)
	for _, v := range strRet {
		intRet = append(intRet, utils.ToInt(v))
	}
	return intRet
}

func getRanges(input []string, i *int) []field {
	arr := make([]field, 0)
	for input[*i] != "" {
		split := strings.Split(input[*i], ": ")
		res := field{name: split[0]}

		split = strings.Split(split[1], " or ")
		r1 := strings.Split(split[0], "-")
		r2 := strings.Split(split[1], "-")

		res.a = utils.ToInt(r1[0])
		res.b = utils.ToInt(r1[1])
		res.c = utils.ToInt(r2[0])
		res.d = utils.ToInt(r2[1])

		arr = append(arr, res)
		*i++
	}
	return arr
}
