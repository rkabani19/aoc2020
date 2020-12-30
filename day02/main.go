package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var data [][]string
	readFile(&data)

	part1(data)
	part2(data)
}

func part2(data [][]string) {
	valid := 0
	for _, val := range data {
		i, _ := strconv.Atoi(val[0])
		j, _ := strconv.Atoi(val[1])
		target := val[2]
		password := val[3]

		iValid := false
		if i > 0 && i <= len(password) {
			if string(password[i-1]) == target {
				iValid = true
			}
		}

		jValid := false
		if j > 0 && j <= len(password) {
			if string(password[j-1]) == target {
				jValid = true
			}
		}

		if (jValid && !iValid) || (!jValid && iValid) {
			valid++
		}
	}

	fmt.Printf("Valid number of passwords for part 2: %d\n", valid)
}

func part1(data [][]string) {
	valid := 0
	for _, val := range data {
		lower, _ := strconv.Atoi(val[0])
		upper, _ := strconv.Atoi(val[1])
		target := val[2]
		password := val[3]
		numOfTarget := strings.Count(password, target)
		if numOfTarget <= upper && numOfTarget >= lower {
			valid++
		}
	}

	fmt.Printf("Valid number of passwords for part 1: %d\n", valid)
}

func readFile(data *[][]string) {
	file, err := os.Open("./input/day2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parsedArr := parseLine(line)
		*data = append(*data, parsedArr)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func parseLine(str string) []string {
	re := regexp.MustCompile(`[^-:\s]+`)
	res := re.FindAllString(str, -1)

	return res
}
