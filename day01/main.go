package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	var data []int
	totalSum := 2020

	readFile(&data)

	result1, err1 := part1(data, totalSum)
	if err1 == nil {
		fmt.Printf("part 1 answer: %d\n", result1)
	}

	result2, err2 := part2(data, totalSum)
	if err2 == nil {
		fmt.Printf("part 2 answer: %d\n", result2)
	}
}

func part1(data []int, totalSum int) (int, error) {
	m := make(map[int]int)

	for i, n := range data {
		idx, ok := m[n]
		if ok {
			return n * data[idx], nil
		}
		m[totalSum-n] = i
	}

	return -1, errors.New("No number found")
}

func part2(data []int, totalSum int) (int, error) {
	n := len(data)
	sort.Ints(data)

	for i := 0; i < n-2; i++ {
		if i > 0 && data[i] == data[i-1] {
			continue
		}
		left := i + 1
		right := n - 1
		target := totalSum - data[i]

		for left < right {
			sum := data[left] + data[right]
			if sum == target {
				return data[i] * data[left] * data[right], nil
			} else if sum > target {
				right--
			} else {
				left++
			}
		}
	}

	return -1, errors.New("No number found")
}

func readFile(data *[]int) {
	file, err := os.Open("day1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		*data = append(*data, i)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
