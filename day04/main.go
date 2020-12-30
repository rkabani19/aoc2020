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

var requiredFields map[string]func(string) bool = map[string]func(string) bool{
	"byr": byrValid,
	"iyr": iyrValid,
	"eyr": eyrValid,
	"hgt": hgtValid,
	"hcl": hclValid,
	"ecl": eclValid,
	"pid": pidValid,
	"cid": cidValid,
}

func main() {
	var data []map[string]string
	readFile(&data)

	a := part1(data)
	fmt.Printf("Part 1: %d\n", a)

	b := part2(data)
	fmt.Printf("Part 2: %d\n", b)
}

func byrValid(str string) bool {
	year, _ := strconv.Atoi(str)
	if year >= 1920 && year <= 2002 {
		return true
	}
	return false
}

func iyrValid(str string) bool {
	year, _ := strconv.Atoi(str)
	if year >= 2010 && year <= 2020 {
		return true
	}
	return false
}

func eyrValid(str string) bool {
	year, _ := strconv.Atoi(str)
	if year >= 2020 && year <= 2030 {
		return true
	}
	return false
}

func hgtValid(str string) bool {
	if strings.HasSuffix(str, "cm") {
		height, _ := strconv.Atoi(strings.Split(str, "cm")[0])
		if height >= 150 && height <= 193 {
			return true
		}
		return false
	} else if strings.HasSuffix(str, "in") {
		height, _ := strconv.Atoi(strings.Split(str, "in")[0])
		if height >= 59 && height <= 76 {
			return true
		}
		return false
	}
	return false
}

func hclValid(str string) bool {
	if strings.HasPrefix(str, "#") {
		color := strings.SplitAfter(str, "#")[1]
		match, _ := regexp.MatchString("[a-f0-9]{6}", color)
		return match
	}
	return false
}

func eclValid(str string) bool {
	if str == "amb" || str == "blu" || str == "brn" || str == "gry" || str == "grn" || str == "hzl" || str == "oth" {
		return true
	}
	return false
}

func pidValid(str string) bool {
	if match, _ := regexp.MatchString("[0-9]{9}", str); match && len(str) == 9 {
		return true
	}
	return false
}

func cidValid(str string) bool {
	return true
}

func part2(data []map[string]string) int {
	validPassports := 0

	for _, m := range data {
		invalid := false
		for k, function := range requiredFields {
			if val, ok := m[k]; (!ok && k != "cid") || !function(val) {
				invalid = true
				break
			}
		}

		if !invalid {
			validPassports++
		}
	}

	return validPassports

}

func part1(data []map[string]string) int {
	validPassports := 0

	for _, m := range data {
		invalid := false
		for k, _ := range requiredFields {
			if _, ok := m[k]; !ok && k != "cid" {
				invalid = true
				break
			}
		}

		if !invalid {
			validPassports++
		}
	}

	return validPassports
}

func readFile(data *[]map[string]string) {
	file, err := os.Open("./input/day4.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	passport := make(map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parsePassport(passport, line)
		if len(line) == 0 {
			*data = append(*data, passport)
			passport = make(map[string]string)
		}
	}
	*data = append(*data, passport)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func parsePassport(passport map[string]string, line string) {
	re := regexp.MustCompile(`[^:\s]+`)
	res := re.FindAllString(line, -1)

	for i := 0; i < len(res); i += 2 {
		key := res[i]
		val := res[i+1]
		passport[key] = val
	}
}
