package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func ReadFile(filepath string) ([]string, error) {
	var data []string
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		data = append(data, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return data, nil
}

func GetMax(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func GetMin(x int, y int) int {
	if x > y {
		return y
	}
	return x
}

func ToInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
	return num
}
