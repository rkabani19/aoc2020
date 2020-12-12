package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var myBag string = "shiny gold"

func main() {
	graph := make(map[string]map[string]int)
	readFile(graph)

	a := part1(graph)
	fmt.Printf("Part 1: %d\n", a)

	b := part2(graph)
	fmt.Printf("Part 2: %d\n", b)
}

func part2(graph map[string]map[string]int) int {
	return countBags(graph, myBag)
}

func countBags(graph map[string]map[string]int, node string) int {
	bags, ok := graph[node]
	if !ok {
		return 0
	}

	count := 0
	for k, v := range bags {
		count += v + v*countBags(graph, k)
	}
	return count
}

func part1(graph map[string]map[string]int) int {
	count := 0
	for m := range graph {
		if containsMyBag(graph, m) {
			count++
		}
	}
	return count
}

func containsMyBag(graph map[string]map[string]int, node string) bool {
	if _, ok := graph[node][myBag]; ok {
		return true
	}

	for bag := range graph[node] {
		if containsMyBag(graph, bag) {
			return true
		}
	}
	return false
}

func readFile(graph map[string]map[string]int) {
	file, err := os.Open("day7.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i := scanner.Text()
		parseLuggage(graph, i)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func parseLuggage(graph map[string]map[string]int, line string) {
	split := strings.Split(line, " bags contain ")
	if split[1] == "no other bags." {
		return
	}
	bags := strings.Split(split[1], ", ")

	parent := split[0]
	graph[parent] = make(map[string]int)

	for _, bag := range bags {
		split := strings.Split(bag, " ")
		num, _ := strconv.Atoi(split[0])
		graph[parent][split[1]+" "+split[2]] = num
	}
}
