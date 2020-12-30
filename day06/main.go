package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var data []map[byte]bool
	readFile(&data)

	a := part1(data)
	fmt.Printf("Part 1: %d\n", a)

	b := part2(data)
	fmt.Printf("Part 2: %d\n", b)
}

func part2(data []map[byte]bool) int {
	count := 0
	for _, m := range data {
		for _, val := range m {
			if val {
				count++
			}
		}
	}
	return count
}

func part1(data []map[byte]bool) int {
	count := 0
	for _, m := range data {
		count += len(m)
	}
	return count
}

func readFile(data *[]map[byte]bool) {
	file, err := os.Open("./input/day6.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	boardingPass := make(map[byte]bool)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			*data = append(*data, boardingPass)
			boardingPass = make(map[byte]bool)
		} else {
			parseBoardingPass(boardingPass, line)
		}
	}
	*data = append(*data, boardingPass)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func parseBoardingPass(m map[byte]bool, line string) {
	mLen := len(m)
	tempM := make(map[byte]bool)

	for i := 0; i < len(line); i++ {
		r := line[i]
		if mLen == 0 {
			m[r] = true
		} else {
			tempM[r] = true
		}
	}

	if mLen == 0 {
		return
	}

	for k := range m {
		_, ok := tempM[k]
		if !ok {
			m[k] = false
		}
	}
}

//Parsing function for part 1
//func parseBoardingPass(m map[byte]bool, line string) {
//	for i := 0; i < len(line); i++ {
//		r := line[i]
//		m[r] = true
//	}
//}
