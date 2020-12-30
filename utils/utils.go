package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func ReadFile(filepath string) ([]string, error) {
	var data []string
	file, err := os.Open("input/" + filepath)
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

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func ReplaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}

func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
