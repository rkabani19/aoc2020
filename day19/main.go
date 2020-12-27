package main

import (
	"fmt"
	"strings"

	"github.com/rkabani19/aoc2020/utils"
)

type Rule struct {
	sub  [][]int
	rule string
}

type message string

func main() {
	input, _ := utils.ReadFile("day19.txt")

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part2(input []string) int {
	rules, messages := parseInput(input)
	rules[8] = Rule{sub: [][]int{{42}, {42, 8}}}
	rules[11] = Rule{sub: [][]int{{42, 31}, {42, 11, 31}}}

	count := 0
	for _, msg := range messages {
		start := 0
		if matches(msg, &start, rules, 0) {
			count++
		}
	}
	return count
}

func part1(input []string) int {
	rules, messages := parseInput(input)
	count := 0
	for _, msg := range messages {
		start := 0
		if matches(msg, &start, rules, 0) {
			count++
		}
	}
	return count
}

func matches(msg message, i *int, rules map[int]Rule, ri int) bool {
	if len(msg) <= *i {
		return false
	}

	rule := rules[ri]
	if rule.rule != "" {
		if string(msg[*i]) == rule.rule {
			return true
		}
		return false
	}

	for _, subRule := range rule.sub {
		if subRule == nil {
			continue
		}

		matched := true
		temp := *i
		for _, r := range subRule {
			if !matches(msg, i, rules, r) {
				matched = false
				break
			}

			*i++
			// Need to come back and generalize loop precense logic
			if *i == len(msg) && (r == 8 || r == 11) {
				break
			}
		}

		if matched && (ri != 0 || (ri == 0 && len(msg) == *i)) {
			*i--
			return true
		} else {
			*i = temp
		}
	}

	return false
}

func parseInput(input []string) (map[int]Rule, []message) {
	rules := make(map[int]Rule)
	i := 0
	for ; input[i] != ""; i++ {
		line := input[i]
		sp := strings.Split(line, ": ")
		ruleNum := utils.ToInt(strings.TrimSpace(sp[0]))
		rule := parseRule(sp[1])
		rules[ruleNum] = rule
	}

	messages := make([]message, 0)
	for i += 1; i < len(input); i++ {
		messages = append(messages, message(input[i]))
	}

	return rules, messages
}

func parseRule(rule string) Rule {
	if rule[0] == '"' {
		return Rule{rule: string(rule[1])}
	}

	sp := strings.Split(rule, " | ")
	res := Rule{sub: make([][]int, 2)}
	for i, r := range sp {
		sp2 := strings.Split(r, " ")
		numRules := make([]int, 0)
		for _, n := range sp2 {
			numRules = append(numRules, utils.ToInt(n))
		}
		res.sub[i] = numRules
	}
	return res
}
