package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var data []string // should be of type byte but I'm too lazy
	readFile(&data)

	// part 1
	a := printNumTrees(data, 3, 1)
	fmt.Printf("Part 1: %d\n", a)

	// part 2
	b := printNumTrees(data, 1, 1)
	c := printNumTrees(data, 5, 1)
	d := printNumTrees(data, 7, 1)
	e := printNumTrees(data, 1, 2)
	fmt.Printf("Part 2: %d\n", a*b*c*d*e)
}

func printNumTrees(data []string, right int, down int) int {
	col := 0
	numTrees := 0

	for i := 0; i < len(data); i += down {
		row := data[i]
		if len(row) <= col {
			col %= len(row)
		}

		cur := string(row[col])
		if cur == "#" {
			numTrees++
		}

		col += right
	}

	return numTrees
}

func readFile(data *[]string) {
	file, err := os.Open("./input/day3.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i := scanner.Text()
		*data = append(*data, i)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
