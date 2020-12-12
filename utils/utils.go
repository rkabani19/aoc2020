package utils

import (
	"bufio"
	"os"
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
